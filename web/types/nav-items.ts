import { LucideIcon } from "lucide-react"

export type NavItem = {
  name: string
  url: string
  icon: LucideIcon
}

export type NavItems = {
  items: NavItem[]
}

export type NavUserProps = {
  user: {
    name: string
    email: string
  }
}