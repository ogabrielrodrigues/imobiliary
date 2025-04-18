"use client"

import { Banknote, BarChart, BookOpenText, Building, ChevronRight, ClipboardCheck, FileText, type LucideIcon, UserRoundCheck, UsersRound } from "lucide-react"

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
import { Plan } from "@/types/plan"
import { usePathname } from "next/navigation"

type NavMainItemType = {
  title: string
  url: string
  icon: LucideIcon
  visibility: 'free' | 'pro'
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
      visibility: 'free',
    },
    {
      title: "Contratos",
      url: "/dashboard/locacao/contratos",
      icon: FileText,
      visibility: 'free',
    },
    {
      title: "Inquilinos",
      url: "/dashboard/locacao/inquilinos",
      icon: UsersRound,
      visibility: 'free',
    },
    {
      title: "Imóveis",
      url: "/dashboard/locacao/imoveis",
      icon: Building,
      visibility: 'free',
    },
    {
      title: "Proprietários",
      url: "/dashboard/locacao/proprietarios",
      icon: UserRoundCheck,
      visibility: 'free',
    },
    {
      title: "Vistorias",
      url: "/dashboard/locacao/vistorias",
      icon: ClipboardCheck,
      visibility: 'pro',
    },
    {
      title: "Relatórios",
      url: "/dashboard/locacao/relatorios",
      icon: BarChart,
      visibility: 'pro',
    },
  ],
}

type NavMainProps = {
  plan: Plan
}

export function NavMain({ plan }: NavMainProps) {
  const pathname = usePathname()

  return (
    <SidebarGroup>
      <SidebarGroupLabel>Gerenciamento</SidebarGroupLabel>
      <SidebarMenu>
        <Collapsible key={item.title} asChild defaultOpen>
          <SidebarMenuItem>
            <SidebarMenuButton asChild>
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
                  {item.items?.filter(item => item.visibility === plan.kind).map((subItem) => (
                    <SidebarMenuSubItem key={subItem.title}>
                      <SidebarMenuSubButton asChild isActive={pathname === subItem.url}>
                        <a href={subItem.url}>
                          <subItem.icon />
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
