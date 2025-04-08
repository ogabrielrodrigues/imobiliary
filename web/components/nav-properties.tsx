"use client"

import { usePathname } from "next/navigation"

import {
  SidebarGroup,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"

import { NavItems } from "@/types/nav-items"
import { cn } from "@/lib/utils"

export function NavProperties({ items }: NavItems) {
  const pathname = usePathname()

  return (
    <SidebarGroup className="group-data-[collapsible=icon]:hidden">
      <SidebarGroupLabel>Imóveis</SidebarGroupLabel>
      <SidebarMenu>
        {items.map((item) => (
          <SidebarMenuItem key={item.name}>
            <a
              href={item.url}
            >
              <SidebarMenuButton tooltip={item.name} variant={
                pathname === item.url ? 'outline' : 'default'
              }>
                <item.icon />
                <span>{item.name}</span>
              </SidebarMenuButton>
            </a>
          </SidebarMenuItem>
        ))}
      </SidebarMenu>
    </SidebarGroup>
  )
}
