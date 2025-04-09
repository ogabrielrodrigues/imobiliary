"use client"

import * as React from "react"

import { NavMain } from "@/components/nav-main"
import { NavUser } from "@/components/nav-user"
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"

import { User } from "@/types/user"
import {
  Banknote,
  BarChart,
  Building,
  ClipboardCheck,
  FileText,
  House,
  HousePlus,
  UserRoundCheck,
  UsersRound
} from "lucide-react"
import Link from "next/link"

const data = {
  navMain: [
    {
      title: "Locação de Imóveis",
      url: "/dashboard/locacao",
      icon: House,
      isActive: true,
      items: [
        {
          title: "Alugueres",
          url: "/dashboard/locacao",
          icon: Banknote,
        },
        {
          title: "Contratos",
          url: "/dashboard/locacao/contratos",
          icon: FileText,
        },
        {
          title: "Inquilinos",
          url: "/dashboard/locacao/inquilinos",
          icon: UsersRound,
        },
        {
          title: "Imóveis",
          url: "/dashboard/locacao/imoveis",
          icon: Building,
        },
        {
          title: "Proprietários",
          url: "/dashboard/locacao/proprietarios",
          icon: UserRoundCheck,
        },
        {
          title: "Vistorias",
          url: "/dashboard/locacao/vistorias",
          icon: ClipboardCheck,
        },
        {
          title: "Relatórios",
          url: "/dashboard/locacao/relatorios",
          icon: BarChart,
        },
      ],
    },
  ],
}

type AppSidebarProps = React.ComponentProps<typeof Sidebar> & {
  user: User | undefined
}

export function AppSidebar({ user, ...props }: AppSidebarProps) {
  return (
    <Sidebar
      className="top-(--header-height) h-[calc(100svh-var(--header-height))]!"
      {...props}
    >
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" className="text-lg font-bold">
              <Link href="/" className="flex items-center gap-2">
                <HousePlus className="!size-5" />
                Imobiliary
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={user} />
      </SidebarFooter>
    </Sidebar>
  )
}
