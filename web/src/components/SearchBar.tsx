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
import { useEffect, useRef, useState } from "react";
import useDebounce from "react-use/lib/useDebounce";
import { useFilterStore } from "@/store/module/filter";
import { useTranslate } from "@/utils/i18n";
import Icon from "./Icon";

const SearchBar = () => {
  const t = useTranslate();
  const filterStore = useFilterStore();
  const [queryText, setQueryText] = useState("");
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    const text = filterStore.getState().text;
    setQueryText(text === undefined ? "" : text);
  }, [filterStore.state.text]);

  useDebounce(
    () => {
      filterStore.setTextFilter(queryText.length === 0 ? undefined : queryText);
    },
    200,
    [queryText]
  );

  const handleTextQueryInput = (e: React.FormEvent<HTMLInputElement>) => {
    const text = e.currentTarget.value;
    setQueryText(text);
  };

  return (
    <div className="w-full h-9 flex-row flex justify-start items-center py-2 px-3 rounded-md bg-gray-200 dark:bg-zinc-700">
      <Icon.Search className="w-4 h-auto opacity-30 dark:text-gray-200" />
      <input
        className="flex ml-2 w-24 grow text-sm outline-none bg-transparent dark:text-gray-200"
        type="text"
        placeholder={t("message.search-placeholder")}
        ref={inputRef}
        value={queryText}
        onChange={handleTextQueryInput}
      />
    </div>
  );
};

export default SearchBar;
