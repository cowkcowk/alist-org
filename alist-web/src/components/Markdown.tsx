// @ts-ignore
import SolidMarkdown from "solid-markdown"
import remarkGfm from "remark-gfm"
import rehypeRaw from "rehype-raw"

export const Markdown = (props: {children?: string }) => {
  return (
    <SolidMarkdown
      remarkPlugins={[remarkGfm]}
      rehypePlugins={[rehypeRaw]}
      children={props.children}
    />
  )
}