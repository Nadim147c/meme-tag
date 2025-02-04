import argparse
import os
import re
import uuid

import easyocr


def rename(path: str, text: str):
    """
    Takes a path and text as inputs and renames the path with text as a basename.
    If the new filename exceeds the OS limit, truncate the text to fit within the limit.
    Only allows alphanumeric characters and commonly used punctuation marks.
    """
    dir_path, old_filename = os.path.split(path)
    _, ext = os.path.splitext(old_filename)

    text = re.sub("[^a-zA-Z0-9_.\\-,'\"\\s]", "", text)

    if not text.strip():
        text = str(uuid.uuid4())

    text = f"[tagged] {text}"

    max_length = os.pathconf(path, "PC_NAME_MAX")
    if len(text + ext) > max_length:
        text = text[: max_length - len(ext)]

    new_path = os.path.join(dir_path, text + ext)

    counter = 1
    while os.path.exists(new_path):
        basename = f"{text} {counter}"
        if len(basename) + len(ext) > max_length:
            text = basename[: max_length - len(ext)]
        new_path = os.path.join(dir_path, basename + ext)
        counter += 1

    print(f"rename: '{path}' -> '{new_path}'")
    os.rename(path, new_path)

    return new_path


def main():
    parser = argparse.ArgumentParser(
        prog="meme-tag",
        description="Set file name to the title of a meme using OCR",
    )

    parser.add_argument("images", nargs="+", help="List of input images to process")

    args = parser.parse_args()

    print("Loading OCR model...")
    reader = easyocr.Reader(["en"], gpu=False)

    for file in args.images:
        if "[tagged]" in file:
            print(f"Already tagged: {file}")
            continue
        words = reader.readtext(file, text_threshold=0.85, detail=0)
        text = " ".join(words)
        rename(file, text)


if __name__ == "__main__":
    main()
