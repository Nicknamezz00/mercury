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
import { FallbackLng, FallbackLngObjList } from "i18next";
import { useTranslation } from "react-i18next";
import i18n, { availableLocales, TLocale } from "@/i18n";
import locales from "@/locales/en.json";
import { NestedKeyOf } from "@/types/utils/nestedKeyOf.types";

export const findLanguageMatch = (codename: string): string => {
  if (codename.length > 2 && availableLocales.includes(codename as TLocale)) {
    return codename as Locale;
  }

  const i18nFallbacks = Object.entries(i18n.store.options.fallbackLng as FallbackLng);
  for (const [main, fallbacks] of i18nFallbacks) {
    if (codename === main) {
      return fallbacks[0] as Locale;
    }
  }

  const shortCode = codename.substring(0, 2);
  if (availableLocales.includes(shortCode as TLocale)) {
    return shortCode as Locale;
  }
  for (const exist of availableLocales) {
    if (shortCode == exist.substring(0, 2)) {
      return exist as Locale;
    }
  }
  return (i18n.store.options.fallbackLng as FallbackLngObjList).default[0] as Locale;
};

type Translations = NestedKeyOf<typeof locales>;
type TypeT = (key: Translations, params?: Record<string, any>) => string;

export const useTranslate = (): TypeT => {
  const { t } = useTranslation<Translations>();
  return t;
};
