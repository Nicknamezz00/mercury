/*
 * MIT License
 *
 * Copyright (c) 2023 Runze Wu
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

INSERT INTO
  message (`id`, `content`, `creator_id`)
VALUES
  (
    1,
    '#Hello 👋 Welcome to mercury.',
    101
  );

INSERT INTO
  message (
  `id`,
  `content`,
  `creator_id`,
  `visibility`
)
VALUES
  (
    2,
    '#TODO
- [x] Take more photos about **🌄 sunset**;
- [x] Clean the room;
- [ ] Read *📖 The Little Prince*;
(👆 click to toggle status)',
    101,
    'PROTECTED'
  );

INSERT INTO
  message (
  `id`,
  `content`,
  `creator_id`,
  `visibility`
)
VALUES
  (
    3,
    '**[Slash](https://github.com/boojack/slash)**: A bookmarking and url shortener, save and share your links very easily.
      ![](https://github.com/boojack/slash/raw/main/resources/demo.gif)

**[TechStack](https://github.com/Get-Tech-Stack/TechStack)**: A browser extension that will display the technology stack of the GitHub repository.
![](https://github.com/Get-Tech-Stack/TechStack/blob/main/img/1.png?raw=true)',
    101,
    'PUBLIC'
  );

INSERT INTO
  message (
  `id`,
  `content`,
  `creator_id`,
  `visibility`
)
VALUES
  (
    4,
    '#TODO
- [x] Take more photos about **🌄 sunset**;
- [ ] Clean the classroom;
- [ ] Watch *👦 The Boys*;
(👆 click to toggle status)
',
    102,
    'PROTECTED'
  );

INSERT INTO
  message (
  `id`,
  `content`,
  `creator_id`,
  `visibility`
)
VALUES
  (
    5,
    '三人行，必有我师焉！👨‍🏫',
    102,
    'PUBLIC'
  );