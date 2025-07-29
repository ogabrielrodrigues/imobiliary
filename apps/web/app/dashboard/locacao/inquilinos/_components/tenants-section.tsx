'use client'

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Tenant } from "@/types/tenant"
import { CircleX, Plus } from "lucide-react"
import Link from "next/link"
import { ChangeEvent, useCallback, useMemo, useState } from "react"
import { TenantsList } from "./tenants-list"

type TenantsSectionProps = {
  tenants: Tenant[]
}

export function TenantsSection({ tenants }: TenantsSectionProps) {
  const [searchTerm, setSearchTerm] = useState<string>("")

  const filtered = useMemo(() => {
    return tenants.filter(tenant => {
      const matchesSearch = searchTerm.trim() === "" ||
        tenant.fullname.toLowerCase().includes(searchTerm.toLowerCase())

      return matchesSearch
    })
  }, [tenants, searchTerm])

  const handleSearch = useCallback((event: ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value)
  }, [])

  const handleClearFilter = useCallback(() => {
    setSearchTerm("")
  }, [])

  const isFilterActive = searchTerm !== ""

  return (
    <section className="flex flex-col gap-6 pb-4">
      <div className="flex items-center justify-between">
        <div className="w-full flex gap-2">
          <Input
            placeholder="Procurar..."
            onChange={handleSearch}
            className="w-4/5 sm:w-1/4"
            value={searchTerm}
          />
          <div className="hidden lg:flex gap-2">
            <Button
              onClick={handleClearFilter}
              disabled={!isFilterActive}
            >
              <CircleX className="size-4 mr-1" />
              Limpar Filtro
            </Button>
          </div>
        </div>

        <Link href="/dashboard/locacao/inquilinos/novo">
          <Button>
            <Plus className="size-4 lg:mr-1" />
            <p className="hidden lg:block">Novo Inquilino</p>
          </Button>
        </Link>
      </div>

      {tenants.length > 0
        ? (<TenantsList tenants={filtered} />)
        : (
          <div className="col-span-full text-center py-8">
            <p className="text-muted-foreground">Nenhum inquilino encontrado.</p>
          </div>
        )}
    </section>
  )
}