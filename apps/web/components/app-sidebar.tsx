"use client"

import * as React from "react"

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem
} from "@/components/ui/sidebar"

import { Manager } from "@/types/manager"
import { House, HousePlus } from "lucide-react"
import Link from "next/link"
import { usePathname } from "next/navigation"
import { NavMain } from "./nav-main"
import { NavManager } from "./nav-manager"

type AppSidebarProps = React.ComponentProps<typeof Sidebar> & {
  manager?: Manager
}

export function AppSidebar({ manager, ...props }: AppSidebarProps) {
  const pathname = usePathname()

  return (
    <Sidebar variant="inset" {...props} >
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" className="text-lg font-bold">
              <Link href="/" className="flex items-center gap-2 text-muted">
                <HousePlus className="!size-5" />
                Imobiliary
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent className="!-space-y-3">
        <SidebarGroup>
          <SidebarMenu>
            <SidebarMenuItem>
              <SidebarMenuButton
                asChild
                isActive={pathname === "/dashboard"}
                className="data-[active=true]:!text-primary data-[active=true]:!bg-ring/20"
              >
                <a href="/dashboard">
                  <House />
                  Dashboard</a>
              </SidebarMenuButton>
            </SidebarMenuItem>
          </SidebarMenu>
        </SidebarGroup>
        <NavMain />
      </SidebarContent>
      <SidebarFooter>
        <NavManager manager={manager} />
      </SidebarFooter>
    </Sidebar>
  )
}
