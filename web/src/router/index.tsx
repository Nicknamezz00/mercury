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
import { lazy } from "react";
import { createBrowserRouter, redirect } from "react-router-dom";
import App from "@/App";
import Home from "@/pages/Home";
import { initialGlobalState, initialUserState } from "@/store/module";

const Root = lazy(() => import("@/layouts/Root"));

const initialGlobalStateLoader = async () => {
  try {
    await initialGlobalState();
  } catch (error) {
    // ignore
  }
  return null;
};

const initialUserStateLoader = async (redirectWhenNotFound = true) => {
  let user = undefined;
  try {
    user = await initialUserState();
  } catch (error) {
    // ignore
  }
  if (!user && redirectWhenNotFound) {
    return redirect("/explore");
  }
  return null;
};

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    loader: () => initialGlobalStateLoader(),
    children: [
      {
        path: "/",
        element: <Root />,
        children: [
          {
            path: "",
            element: <Home />,
            loader: () => initialUserStateLoader(),
          },
        ],
      },
    ],
  },
]);

export default router;
