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
import SearchBar from "@/components/SearchBar";
import TagList from "@/components/TagList";
import UsageHeatMap from "@/components/UsageHeatMap";
import { useLayoutStore } from "@/store/module";

const HomeSidebar = () => {
  const layoutStore = useLayoutStore();
  const showHomeSidebar = layoutStore.state.showHomeSidebar;

  return (
    <div
      className={`fixed md:sticky top-0 left-0 w-full md:w-56 h-full shrink-0 pointer-events-none md:pointer-events-auto z10 ${
        showHomeSidebar && "pointer-events-auto"
      }`}
    >
      <div
        className={`fixed top-0 left-0 w-full h-full bg-black opacity-0 pointer-events-none transition-opacity duration-300 md:!hidden ${
          showHomeSidebar && "opacity-60 pointer-events-auto"
        }`}
        onClick={() => layoutStore.setHomeSidebarStatus(false)}
      ></div>
      <aside
        className={`absolute md:relative top-0 right-0 w-56 pt-2 md:w-full h-full max-h-screen overflow-auto hide-scrollbar flex flex-col justify-start items-start py-4 z-30 bg-zinc-100 dark:bg-zinc-800 md:bg-transparent md:shadow-none transition-all duration-300 translate-x-full md:translate-x-0 ${
          showHomeSidebar && "!translate-x-0 shadow-2xl"
        }`}
      >
        <div className="px-4 pr-8 mb-4 w-full">
          <SearchBar />
        </div>
        <UsageHeatMap />
        <TagList />
      </aside>
    </div>
  );
};

export default HomeSidebar;
