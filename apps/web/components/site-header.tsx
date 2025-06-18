"use client"

import { Sidebar } from "lucide-react"
import React from 'react'

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb"
import { Button } from "@/components/ui/button"
import { Separator } from "@/components/ui/separator"
import { useSidebar } from "@/components/ui/sidebar"
import { generateBreadcrumbPath, routes } from "@/lib/routes"
import { usePathname } from "next/navigation"

export function SiteHeader(): React.ReactElement {
  const { toggleSidebar } = useSidebar();
  const pathname = usePathname();

  const breadcrumbPath = generateBreadcrumbPath({
    pathname,
    routeStructure: routes
  });

  return (
    <header className="bg-background sticky top-0 z-50 flex w-full items-center border-b">
      <div className="flex h-(--header-height) w-full items-center gap-2 pl-4 pr-8">
        <Button
          className="h-8 w-8 text-muted-foreground"
          variant="ghost"
          size="icon"
          onClick={toggleSidebar}
        >
          <Sidebar />
        </Button>
        <Separator orientation="vertical" className="mr-2 h-4" />
        <Breadcrumb className="hidden sm:block">
          <BreadcrumbList>
            {breadcrumbPath.map((item, index) => (
              <React.Fragment key={`breadcrumb-${item.href}-${index}`}>
                <BreadcrumbItem>
                  {item.isLast ? (
                    <BreadcrumbPage>{item.name}</BreadcrumbPage>
                  ) : (
                    <BreadcrumbLink href={item.href}>{item.name}</BreadcrumbLink>
                  )}
                </BreadcrumbItem>
                {!item.isLast && <BreadcrumbSeparator />}
              </React.Fragment>
            ))}
          </BreadcrumbList>
        </Breadcrumb>
        {/* <SearchForm className="w-full sm:ml-auto sm:w-1/6" /> */}
      </div>
    </header>
  );
}