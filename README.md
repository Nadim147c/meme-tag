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

- [Python3](https://www.python.org/)
- [uv](https://docs.astral.sh/uv/) (Optional if you know how to manage virtual env)

## 🚀 Installation

**Clone the git repository**

```
git clone --branch easyocr --single-branch https://github.com/Nadim147c/meme-tag.git
cd meme-tag
```

**Install dependencies**

> Note: By default it uses the CPU version of PyTorch. Which downloads fast but runs on CPU (slow).
> If you have a nvidia GPU, you should install the GPU version of PyTorch by removing the following
> lines from `./pyproject.toml` and then run `uv install`

```
[tool.uv.sources]
torch = { index = "pytorch_cpu" }
torchvision = { index = "pytorch_cpu" }
```

```
uv install
```

If you want to installed `meme-tag` command to your system PATH. Then,

```
uv tool install .
```

## 📸 Usage

```sh
# Rename a single image
uv run main.py /path/to/image
# If you have installed it to your system PATH,
meme-tag /path/to/image

# Process all jpg files a directory
uv run main.py /path/to/*.jpg
# If you have installed it to your system PATH,
meme-tag /path/to/*.jpg
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
