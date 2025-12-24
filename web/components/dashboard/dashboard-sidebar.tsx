"use client"

import { IconHome, IconHomePlus } from "@tabler/icons-react"
import Link from "next/link"
import { usePathname } from "next/navigation"
import type { ComponentProps } from "react"
import { ManagerSection } from "@/components/dashboard/manager-section"
import { Navigation } from "@/components/dashboard/navigation"
import {
	Sidebar,
	SidebarContent,
	SidebarFooter,
	SidebarGroup,
	SidebarHeader,
	SidebarMenu,
	SidebarMenuButton,
	SidebarMenuItem,
} from "@/components/ui/sidebar"
import type { Manager } from "@/types/manager"

interface AppSidebarProps extends ComponentProps<typeof Sidebar> {
	manager?: Manager
}

export function DashboardSidebar({ manager, ...props }: AppSidebarProps) {
	const pathname = usePathname()

	return (
		<Sidebar variant="inset" {...props}>
			<SidebarHeader>
				<SidebarMenu>
					<SidebarMenuItem>
						<SidebarMenuButton className="text-lg font-bold">
							<Link
								className="flex items-center gap-2 text-muted text-xl"
								href="/"
							>
								<IconHomePlus className="size-5" />
								imobiliary
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
								className="data-[active=true]:!text-primary data-[active=true]:!bg-ring/20"
								isActive={pathname === "/dashboard"}
							>
								<a href="/dashboard">
									<IconHome />
									<span>Dashboard</span>
								</a>
							</SidebarMenuButton>
						</SidebarMenuItem>
					</SidebarMenu>
				</SidebarGroup>

				<Navigation />
			</SidebarContent>

			<SidebarFooter>
				<ManagerSection manager={manager} />
			</SidebarFooter>
		</Sidebar>
	)
}
