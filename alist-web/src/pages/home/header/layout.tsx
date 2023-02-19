import { Menu, MenuContent, MenuItem } from "@hope-ui/solid"
import { BsGridFill, BsCardImage } from "solid-icons/bs"
import { FaSolidListUl } from "solid-icons/fa"
import { For } from "solid-js"
import { Dynamic } from "solid-js/web"

const layouts = {
  list: FaSolidListUl,
  grid: BsGridFill,
  image: BsCardImage,
} as const

export const Layout = () => {
  return (
    <Menu>
      <MenuContent>
        <For each={Object.entries(layouts)}>
          {(item) => (
            <MenuItem
              icon={<Dynamic component={item[1]} />}
              onSelect={() => {

              }}
            >
              hi
            </MenuItem>
          )}
        </For>
      </MenuContent>
    </Menu>
  )
}