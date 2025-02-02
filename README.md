# meme-tag

### Rename your images based on their **true identity** (as seen by AI)

> Ever looked at your chaotic `Downloads/` folder and wondered, _"What in the name of JPEG is `IMG_42069.png`?"_ Well, wonder no more!
>
> 🚀 **meme-tag** uses **OCR magic** to peer into your images, extract any visible text, and rename them accordingly. Now, your memes, screenshots, and cursed images will finally have names that make sense (or at least, names that an AI thinks make sense).

## 💀 Warnings

- If OCR reads something _incorrectly_, that's on you. Or on the AI. Or on the universe. Either way, we accept no responsibility.
- Meme-tag might rename **every** image in your folder, so maybe don’t run it in `~/Desktop/` unless you like surprises.
- Some images are just... beyond comprehension. If it renames a file to `"???".png`, that's your cue to rethink your life choices.

## 🛠️ Why?

Because naming files manually is **boring**. Let the AI do the work, and **embrace the chaos**.

## 🛠 Dependencies

This tool relies on Tesseract-OCR. You need to install both the library and headers:

#### Linux (ArchLinux) [I use arch BTW]

```sh
pacman -S tesseract tesseract-data-ocr tesseract-data-eng
```

#### Linux (Debian/Ubuntu)

```sh
sudo apt install tesseract-ocr libtesseract-dev
```

#### MacOS (Homebrew)

```sh
brew install tesseract
```

#### Windows

Download and install Tesseract from: https://github.com/tesseract-ocr/tesseract

Ensure tesseract is in your system's PATH.

## 🚀 Installation

```
go install github.com/Nadim147c/meme-tag@latest
```

Or, you know, build it yourself like a responsible adult:

```
git clone https://github.com/Nadim147c/meme-tag.git
cd meme-tag
go install
```

## 📸 Usage

```sh
# Rename a single image
meme-tag my_meme.png

# Process an entire directory
meme-tag -d ~/Pictures
```

Example:
Before:

```
IMG_1337.jpg
```

After:

```
"We live in a society.jpg"
```

Because filenames should tell a **story**.

## 🎉 Contribute

Want to improve this mess? PRs welcome! Just make sure your commits are funny.

🐸 **meme-tag** – The only CLI that gives your memes the names they deserve.
