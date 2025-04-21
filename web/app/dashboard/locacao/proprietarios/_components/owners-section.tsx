'use client'

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Owner } from "@/types/owner"
import { CircleX, Plus } from "lucide-react"
import Link from "next/link"
import { ChangeEvent, useCallback, useMemo, useState } from "react"
import { OwnersList } from "./owners-list"
// import { PropertyList } from "./properties-list"

type OwnersSectionProps = {
  owners: Owner[]
}

export function OwnersSection({ owners }: OwnersSectionProps) {
  const [searchTerm, setSearchTerm] = useState<string>("")
  const [selectedKind, setSelectedKind] = useState<string>("Todos")
  const [selectedStatus, setSelectedStatus] = useState<string>("Todos")

  const filtered = useMemo(() => {
    return owners.filter(owner => {
      const matchesSearch = searchTerm.trim() === "" ||
        owner.address.mini_address.toLowerCase().includes(searchTerm.toLowerCase())

      return matchesSearch
    })
  }, [owners, searchTerm])

  const handleSearch = useCallback((event: ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value)
  }, [])

  const handleClearFilter = useCallback(() => {
    setSearchTerm("")
  }, [])

  const isFilterActive = searchTerm !== ""

  return (
    <section className="flex flex-col gap-6">
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
              Limpar Filtros
            </Button>
          </div>
        </div>

        <Link href="/dashboard/locacao/proprietarios/novo">
          <Button>
            <Plus className="size-4 mr-1" />
            <p className="hidden lg:block">Novo Proprietário</p>
          </Button>
        </Link>
      </div>

      <OwnersList owners={filtered} />
    </section>
  )
}