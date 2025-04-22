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

import { Plan } from "@/types/plan"
import { User } from "@/types/user"
import { House, HousePlus, Sparkles } from "lucide-react"
import Link from "next/link"
import { usePathname } from "next/navigation"
import { NavMain } from "./nav-main"
import { NavUser } from "./nav-user"

type AppSidebarProps = React.ComponentProps<typeof Sidebar> & {
  user: User | undefined
  plan: Plan
}

export function AppSidebar({ user, plan, ...props }: AppSidebarProps) {
  const pathname = usePathname()

  return (
    <Sidebar
      className="top-(--header-height) h-[calc(100svh-var(--header-height))]! "
      {...props}
    >
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" className="text-lg font-bold">
              <Link href="/" className="flex items-center gap-2">
                <HousePlus className="!size-5" />
                Imobiliary
                {plan.kind === 'pro' && <Sparkles className="size-4 self-start" />}
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent className="!-space-y-3">
        <SidebarGroup>
          <SidebarMenu>
            <SidebarMenuItem>
              <SidebarMenuButton asChild isActive={pathname === "/dashboard"}>
                <Link href="/dashboard">
                  <House />
                  Dashboard</Link>
              </SidebarMenuButton>
            </SidebarMenuItem>
          </SidebarMenu>
        </SidebarGroup>
        <NavMain plan={plan} />
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={user} plan={plan} />
      </SidebarFooter>
    </Sidebar>
  )
}
