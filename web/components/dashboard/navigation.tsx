"use client"

import {
	type Icon,
	IconHomeHand,
	IconLibrary,
	IconUserStar,
	IconUsers,
} from "@tabler/icons-react"
import { ChevronRight, type LucideIcon } from "lucide-react"
import { usePathname } from "next/navigation"
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

interface NavigationItem {
	title: string
	href: string
	icon: LucideIcon
}

interface NavigationRoute {
	title: string
	href: string
	icon: Icon
	items: NavigationItem[]
}

const navigation: NavigationRoute = {
	title: "Locação",
	href: "/dashboard/locacao",
	icon: IconLibrary,
	items: [
		{
			title: "Inquilinos",
			href: "/dashboard/locacao/inquilinos",
			icon: IconUsers,
		},
		{
			title: "Imóveis",
			href: "/dashboard/locacao/imoveis",
			icon: IconHomeHand,
		},
		{
			title: "Proprietários",
			href: "/dashboard/locacao/proprietarios",
			icon: IconUserStar,
		},
	],
}

export function Navigation() {
	const pathname = usePathname()

	return (
		<SidebarGroup>
			<SidebarGroupLabel>Gerenciamento</SidebarGroupLabel>
			<SidebarMenu>
				<Collapsible asChild defaultOpen key={navigation.title}>
					<SidebarMenuItem>
						<SidebarMenuButton
							asChild
							className="data-[active=true]:!text-primary data-[active=true]:!bg-ring/20 data-[active=true]:!font-medium"
							isActive={pathname === "/dashboard/locacao"}
						>
							<a href={navigation.href}>
								<navigation.icon />
								<span>{navigation.title}</span>
							</a>
						</SidebarMenuButton>

						<CollapsibleTrigger asChild>
							<SidebarMenuAction className="data-[state=open]:rotate-90">
								<ChevronRight />
								<span className="sr-only">Toggle</span>
							</SidebarMenuAction>
						</CollapsibleTrigger>

						<CollapsibleContent>
							<SidebarMenuSub>
								{navigation.items?.map(item => (
									<SidebarMenuSubItem key={item.title}>
										<SidebarMenuSubButton
											asChild
											className="group data-[active=true]:!text-primary data-[active=true]:!bg-ring/20 data-[active=true]:!font-medium"
											isActive={pathname === item.href}
										>
											<a href={item.href}>
												<item.icon className="group-data-[active=true]:!text-primary" />
												<span>{item.title}</span>
											</a>
										</SidebarMenuSubButton>
									</SidebarMenuSubItem>
								))}
							</SidebarMenuSub>
						</CollapsibleContent>
					</SidebarMenuItem>
				</Collapsible>
			</SidebarMenu>
		</SidebarGroup>
	)
}
