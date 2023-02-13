import { createResource } from "solid-js"
import { Markdown } from "~/components"

const fetchReadme = async () =>
  await (
    await fetch("https://jsd.nn.ci/gh/alist-org/alist@main/README.md")
  ).text()

const About = () => {
  const [readme] = createResource(fetchReadme)
  return (
    <Markdown children={readme() } />
  )
}

export default About
