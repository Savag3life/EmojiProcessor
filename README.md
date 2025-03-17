# Discord PNG-Emoji Splitter
This is a simple project which when ran, starts a small webapp to upload a full-sized PNG, along with a descriptor & emoji size to generate a set of emoji's matching in the input image. 

## How to run:
```bash
git clone https://github.com/Savag3life/EmojiProcessor.git
cd EmojiProcessor
go run main.go
```
The default port for the webserver is hosted on `localhost:8080`.

## Example
With the following input image,
![Example](https://github.com/Savag3life/EmojiProcessor/blob/main/example/flow_eyes.png)

We generate the following emoji's with a descriptor or `eyes` and a size of `64x64` pixels per-emote.

```aiignore
:eyes_0_0::eyes_0_1::eyes_0_2::eyes_0_3::eyes_0_4::eyes_0_5::eyes_0_6::eyes_0_7::eyes_0_8::eyes_0_9:
:eyes_1_0::eyes_1_1::eyes_1_2::eyes_1_3::eyes_1_4::eyes_1_5::eyes_1_6::eyes_1_7::eyes_1_8::eyes_1_9:
```

Which creates 20 individual emoji's, each with a size of `64x64` pixels. You can use any of them in any order, or all of them.

```:eyes_1_1::eyes_1_2::eyes_1_3::eyes_1_7::eyes_1_8::eyes_1_9:```

![Example](https://github.com/Savag3life/EmojiProcessor/blob/main/example/output/eyes_1_1.png)![Example](https://github.com/Savag3life/EmojiProcessor/blob/main/example/output/eyes_1_2.png)![Example](https://github.com/Savag3life/EmojiProcessor/blob/main/example/output/eyes_1_3.png)![Example](https://github.com/Savag3life/EmojiProcessor/blob/main/example/output/eyes_1_7.png)![Example](https://github.com/Savag3life/EmojiProcessor/blob/main/example/output/eyes_1_8.png)![Example](https://github.com/Savag3life/EmojiProcessor/blob/main/example/output/eyes_1_9.png)