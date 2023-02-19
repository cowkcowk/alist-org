
const settings: Record<string, string> = {}

export const setSettings = (items: Record<string, string>) => {
  Object.keys(items)
}

export const getSetting = (key: string) => settings[key] ?? ""

export const getMainColor = (): string => {
  if (window.ALIST.main_color) {
    return window.ALIST.main_color
  }
  return getSetting("main_color") || "#1890ff"
}