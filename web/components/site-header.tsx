"use client"

import { Sidebar } from "lucide-react"

import { SearchForm } from "@/components/search-form"
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
import { usePathname } from "next/navigation"

type PageType = {
  name: string;
  href: string;
  menu?: {
    name: string;
    href: string;
  };
}

const pages: Record<string, PageType> = {
  "/dashboard": {
    name: "Visão Geral",
    href: "/dashboard",
  },
  "/dashboard/conta": {
    name: "Conta",
    href: "/dashboard/conta",
  },
  "/dashboard/assinatura": {
    name: "Assinatura",
    href: "/dashboard/assinatura",
  },
  "/dashboard/locacao": {
    name: "Alugueis",
    href: "/dashboard/locacao",
  },
  "/dashboard/locacao/contratos": {
    name: "Contratos",
    href: "/dashboard/locacao/contratos",
    menu: {
      name: "Locação",
      href: "/dashboard/locacao",
    }
  },
  "/dashboard/locacao/inquilinos": {
    name: "Inquilinos",
    href: "/dashboard/locacao/inquilinos",
    menu: {
      name: "Locação",
      href: "/dashboard/locacao",
    }
  },
  "/dashboard/locacao/imoveis": {
    name: "Imóveis",
    href: "/dashboard/locacao/imoveis",
    menu: {
      name: "Locação",
      href: "/dashboard/locacao",
    }
  },
  "/dashboard/locacao/proprietarios": {
    name: "Proprietários",
    href: "/dashboard/locacao/proprietarios",
    menu: {
      name: "Locação",
      href: "/dashboard/locacao",
    }
  },
  "/dashboard/locacao/vistorias": {
    name: "Vistorias",
    href: "/dashboard/locacao/vistorias",
    menu: {
      name: "Locação",
      href: "/dashboard/locacao",
    }
  },
  "/dashboard/locacao/relatorios": {
    name: "Relatórios",
    href: "/dashboard/locacao/relatorios",
    menu: {
      name: "Locação",
      href: "/dashboard/locacao",
    }
  }
}
export function SiteHeader() {
  const { toggleSidebar } = useSidebar()
  const pathname = usePathname()
  const currentPage = pathname ? pages[pathname] : undefined

  return (
    <header className="bg-background sticky top-0 z-50 flex w-full items-center border-b">
      <div className="flex h-(--header-height) w-full items-center gap-2 pl-4 pr-8">
        <Button
          className="h-8 w-8"
          variant="ghost"
          size="icon"
          onClick={toggleSidebar}
        >
          <Sidebar />
        </Button>
        <Separator orientation="vertical" className="mr-2 h-4" />
        <Breadcrumb className="hidden sm:block">
          <BreadcrumbList>
            <BreadcrumbItem>
              <BreadcrumbLink href="/dashboard">
                Dashboard
              </BreadcrumbLink>
            </BreadcrumbItem>
            <BreadcrumbSeparator />
            {currentPage?.menu ? (
              <>
                <BreadcrumbItem>
                  <BreadcrumbLink href={currentPage.menu.href}>
                    {currentPage.menu.name}
                  </BreadcrumbLink>
                </BreadcrumbItem>
                <BreadcrumbSeparator />
                <BreadcrumbItem>
                  <BreadcrumbPage>{currentPage.name}</BreadcrumbPage>
                </BreadcrumbItem>
              </>
            ) : (
              <BreadcrumbItem>
                <BreadcrumbPage>{currentPage?.name}</BreadcrumbPage>
              </BreadcrumbItem>
            )}
          </BreadcrumbList>
        </Breadcrumb>
        <SearchForm className="w-full sm:ml-auto sm:w-1/6" />
      </div>
    </header>
  )
}
