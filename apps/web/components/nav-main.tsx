"use client"

import { Banknote, BookOpenText, Building, ChevronRight, type LucideIcon, UserRoundCheck, UsersRound } from "lucide-react"

import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible"
import {
  SidebarGroup,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuAction,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
} from "@/components/ui/sidebar"
import { usePathname } from "next/navigation"

type NavMainItemType = {
  title: string
  url: string
  icon: LucideIcon
}

type NavMainType = {
  title: string
  url: string
  icon: LucideIcon
  items: NavMainItemType[]
}

const item: NavMainType = {
  title: "Locação de Imóveis",
  url: "/dashboard/locacao",
  icon: BookOpenText,
  items: [
    {
      title: "Alugueres",
      url: "/dashboard/locacao",
      icon: Banknote,
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
  ],
}

export function NavMain() {
  const pathname = usePathname()

  return (
    <SidebarGroup>
      <SidebarGroupLabel>Gerenciamento</SidebarGroupLabel>
      <SidebarMenu>
        <Collapsible key={item.title} asChild defaultOpen>
          <SidebarMenuItem>
            <SidebarMenuButton
              asChild
              isActive={item.url.includes(pathname)}
              className="data-[active=true]:!text-primary data-[active=true]:!bg-sidebar"
            >
              <a href={item.url}>
                <item.icon />
                <span>{item.title}</span>
              </a>
            </SidebarMenuButton>
            <>
              <CollapsibleTrigger asChild>
                <SidebarMenuAction className="data-[state=open]:rotate-90">
                  <ChevronRight />
                  <span className="sr-only">Toggle</span>
                </SidebarMenuAction>
              </CollapsibleTrigger>
              <CollapsibleContent>
                <SidebarMenuSub>
                  {item.items?.map((subItem) => (
                    <SidebarMenuSubItem key={subItem.title}>
                      <SidebarMenuSubButton
                        asChild
                        isActive={pathname === subItem.url}
                        className="group data-[active=true]:!text-primary data-[active=true]:!bg-ring/20 data-[active=true]:!font-medium"
                      >
                        <a href={subItem.url}>
                          <subItem.icon className="group-data-[active=true]:!text-primary" />
                          <span>{subItem.title}</span>
                        </a>
                      </SidebarMenuSubButton>
                    </SidebarMenuSubItem>
                  ))}
                </SidebarMenuSub>
              </CollapsibleContent>
            </>
          </SidebarMenuItem>
        </Collapsible>
      </SidebarMenu>
    </SidebarGroup>
  )
}
