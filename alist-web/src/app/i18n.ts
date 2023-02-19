import { createI18nContext } from "@solid-primitives/i18n"
import { createSignal } from "solid-js"

export let initialLang = localStorage.getItem("lang") ?? ""

const i18n = createI18nContext({}, initialLang)

export { i18n }