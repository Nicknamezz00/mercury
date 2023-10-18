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
import { Outlet } from "react-router-dom";

function Root() {
  return (
    <div className="w-full min-h-full bg-zinc-100 dark:bg-zinc-800">
      <div className="w-full h-auto flex flex-col justify-start items-center">
        <div>ROOT</div>
      </div>
      <div className="w-full max-w-6xl mx-auto flex flex-row justify-center items-start sm:px-4">
        <div>HEADER</div>
        <main className="w-auto max-w-full sm:max-w-[calc(100%-14rem)] flex-grow shrink flex flex-col justify-start items-start">
          <Outlet />
        </main>
      </div>
    </div>
  );
}

export default Root;
