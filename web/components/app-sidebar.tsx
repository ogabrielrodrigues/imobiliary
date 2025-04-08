"use client"

import * as React from "react"

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"

import { ChartPie, Eye, HandCoins, House, HousePlus, MapPinHouse, Scroll } from "lucide-react"

import { NavProperties } from "@/components/nav-properties"
import { NavMain } from "@/components/nav-main"
import { NavUser } from "@/components/nav-user"

// Mock data
const data = {
  user: {
    name: "Gabriel Rodrigues",
    email: "gabriel.rodrigues@creci.org.br",
  },
  navMain: [
    {
      name: "Visão Geral",
      url: "/dashboard",
      icon: Eye,
    },
    {
      name: "Relatórios (SAAEC)",
      url: "/dashboard/relatorios",
      icon: Scroll,
    },
  ],
  properties: [
    {
      name: "Compra/Venda",
      url: "/dashboard/imoveis/compra-venda",
      icon: HandCoins,
    },
    {
      name: "Alugueis",
      url: "/dashboard/imoveis/alugueis",
      icon: MapPinHouse,
    },
  ],
}

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  return (
    <Sidebar collapsible="offcanvas" {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton
              asChild
              className="data-[slot=sidebar-menu-button]:!p-1.5"
            >
              <a href="/dashboard">
                <HousePlus className="!size-5" />
                <span className="text-base font-semibold">Imobiliary</span>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
        {/* <NavProperties items={data.properties} /> */}
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={data.user} />
      </SidebarFooter>
    </Sidebar>
  )
}
