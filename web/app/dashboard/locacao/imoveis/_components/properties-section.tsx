'use client'

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Property } from "@/types/property"
import { CircleCheck, CircleMinus, CircleX, Hammer, LockKeyhole, Plus } from "lucide-react"
import Link from "next/link"
import { ChangeEvent, useCallback, useMemo, useState } from "react"
import { PropertyList } from "./properties-list"

type PropertiesSectionProps = {
  properties: Property[]
}

export function PropertiesSection({ properties }: PropertiesSectionProps) {
  const [searchTerm, setSearchTerm] = useState<string>("")
  const [selectedKind, setSelectedKind] = useState<string>("Todos")
  const [selectedStatus, setSelectedStatus] = useState<string>("Todos")

  const filtered = useMemo(() => {
    return properties.filter(property => {
      const matchesSearch = searchTerm.trim() === "" ||
        property.address.address.toLowerCase().includes(searchTerm.toLowerCase())

      const matchesKind = selectedKind === "Todos" ||
        property.kind === selectedKind

      const matchesStatus = selectedStatus === "Todos" ||
        property.status === selectedStatus

      return matchesSearch && matchesKind && matchesStatus
    })
  }, [properties, searchTerm, selectedKind, selectedStatus])

  const handleSearch = useCallback((event: ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value)
  }, [])

  const handleFilterKind = useCallback((value: string) => {
    setSelectedKind(value)
  }, [])

  const handleFilterStatus = useCallback((value: string) => {
    setSelectedStatus(value)
  }, [])

  const handleClearFilter = useCallback(() => {
    setSearchTerm("")
    setSelectedKind("Todos")
    setSelectedStatus("Todos")
  }, [])

  const isFilterActive = searchTerm !== "" || selectedKind !== "Todos" || selectedStatus !== "Todos"

  return (
    <section className="flex flex-col gap-6">
      <div className="flex items-center justify-between">
        <div className="w-full flex flex-col sm:flex-row gap-2">
          <Input
            placeholder="Procurar..."
            onChange={handleSearch}
            className="w-4/5 sm:w-3/6"
            value={searchTerm}
          />
          <div className="hidden flex-col sm:flex-row gap-2 sm:flex">
            <Select onValueChange={handleFilterKind} value={selectedKind}>
              <SelectTrigger>
                <SelectValue placeholder="Tipo" />
              </SelectTrigger>
              <SelectContent className="text-sm md:text-base">
                <SelectItem value="Todos">Todos</SelectItem>
                <SelectItem value="Residencial">Residencial</SelectItem>
                <SelectItem value="Comercial">Comercial</SelectItem>
                <SelectItem value="Industrial">Industrial</SelectItem>
                <SelectItem value="Terreno">Terreno</SelectItem>
                <SelectItem value="Rural">Rural</SelectItem>
              </SelectContent>
            </Select>

            <Select onValueChange={handleFilterStatus} value={selectedStatus}>
              <SelectTrigger>
                <SelectValue placeholder="Status" />
              </SelectTrigger>
              <SelectContent className="text-sm md:text-base">
                <SelectItem value="Todos">Todos</SelectItem>
                <SelectItem value="Disponível">
                  <div className="flex items-center gap-2">
                    <CircleCheck className="size-4 text-emerald-500" />
                    <span>Disponível</span>
                  </div>
                </SelectItem>
                <SelectItem value="Ocupado">
                  <div className="flex items-center gap-2">
                    <CircleMinus className="size-4 text-yellow-500" />
                    <span>Ocupado</span>
                  </div>
                </SelectItem>
                <SelectItem value="Indisponível">
                  <div className="flex items-center gap-2">
                    <CircleX className="size-4 text-red-500" />
                    <span>Indisponível</span>
                  </div>
                </SelectItem>
                <SelectItem value="Reservado">
                  <div className="flex items-center gap-2">
                    <LockKeyhole className="size-4 text-sky-500" />
                    <span>Reservado</span>
                  </div>
                </SelectItem>
                <SelectItem value="Reformando">
                  <div className="flex items-center gap-2">
                    <Hammer className="size-4 text-orange-500" />
                    <span>Reformando</span>
                  </div>
                </SelectItem>
              </SelectContent>
            </Select>

            <Button
              onClick={handleClearFilter}
              disabled={!isFilterActive}
            >
              <CircleX className="size-4 lg:mr-1" />
              <p className="hidden lg:block">Limpar Filtros</p>
            </Button>
          </div>
        </div>

        <Link href="/dashboard/locacao/imoveis/novo">
          <Button>
            <Plus className="size-4 lg:mr-1" />
            <p className="hidden lg:block">Novo Imóvel</p>
          </Button>
        </Link>
      </div>

      <PropertyList properties={filtered} />
    </section>
  )
}