"use client"

import React from 'react'

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb"
import { Separator } from "@/components/ui/separator"
import { SidebarTrigger, useSidebar } from "@/components/ui/sidebar"
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
    <header className="flex h-16 shrink-0 items-center gap-2">
      <div className="flex items-center gap-2 px-4">
        <SidebarTrigger className="-ml-1" />
        <Separator
          orientation="vertical"
          className="mr-2 data-[orientation=vertical]:h-4"
        />
        <Breadcrumb>
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
      </div>
    </header>

    // <header className="bg-background absolute top-0 left-0 z-50 flex w-full items-center border-b">
    //   <div className="flex h-(--header-height) w-full items-center gap-2 pl-4 pr-8">
    //     <Button
    //       className="h-8 w-8 text-muted-foreground"
    //       variant="ghost"
    //       size="icon"
    //       onClick={toggleSidebar}
    //     >
    //       <Sidebar />
    //     </Button>
    //     <Separator orientation="vertical" className="mr-2 h-4" />
    //     <Breadcrumb className="hidden sm:block">
    //       <BreadcrumbList>
    //         {breadcrumbPath.map((item, index) => (
    //           <React.Fragment key={`breadcrumb-${item.href}-${index}`}>
    //             <BreadcrumbItem>
    //               {item.isLast ? (
    //                 <BreadcrumbPage>{item.name}</BreadcrumbPage>
    //               ) : (
    //                 <BreadcrumbLink href={item.href}>{item.name}</BreadcrumbLink>
    //               )}
    //             </BreadcrumbItem>
    //             {!item.isLast && <BreadcrumbSeparator />}
    //           </React.Fragment>
    //         ))}
    //       </BreadcrumbList>
    //     </Breadcrumb>
    //     {/* <SearchForm className="w-full sm:ml-auto sm:w-1/6" /> */}
    //   </div>
    // </header>
  );
}