"use client"

import { usePathname } from "next/navigation"
import React from "react"
import {
	Breadcrumb,
	BreadcrumbItem,
	BreadcrumbLink,
	BreadcrumbList,
	BreadcrumbPage,
	BreadcrumbSeparator,
} from "@/components/ui/breadcrumb"
import { Separator } from "@/components/ui/separator"
import { SidebarTrigger } from "@/components/ui/sidebar"
import { generateBreadcrumbPath, routes } from "@/lib/routes"

export function DashboardHeader() {
	const pathname = usePathname()

	const breadcrumbPath = generateBreadcrumbPath({
		pathname,
		routeStructure: routes,
	})

	return (
		<header className="flex h-16 shrink-0 items-center gap-2">
			<div className="flex items-center gap-2 px-4">
				<SidebarTrigger className="-ml-1" />

				<Separator
					className="mr-2 data-[orientation=vertical]:h-4"
					orientation="vertical"
				/>

				<Breadcrumb>
					<BreadcrumbList>
						{breadcrumbPath.map((item, index) => (
							<React.Fragment key={`breadcrumb-${item.href}-${index}`}>
								<BreadcrumbItem>
									{item.isLast ? (
										<BreadcrumbPage>{item.name}</BreadcrumbPage>
									) : (
										<BreadcrumbLink href={item.href}>
											{item.name}
										</BreadcrumbLink>
									)}
								</BreadcrumbItem>
								{!item.isLast && <BreadcrumbSeparator />}
							</React.Fragment>
						))}
					</BreadcrumbList>
				</Breadcrumb>
			</div>
		</header>
	)
}
