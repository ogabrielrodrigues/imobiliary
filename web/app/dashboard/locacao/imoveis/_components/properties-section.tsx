'use client'

import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Card, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { cn } from "@/lib/utils"
import { Property } from "@/types/property"
import { ArrowUpRight, CircleCheck, CircleMinus, CircleX, Hammer, LockKeyhole, Plus } from "lucide-react"
import Link from "next/link"
import { ChangeEvent, useState } from "react"
import { bgColorStatusDetail, StatusBadge } from "./status-badge"

type PropertiesSectionProps = {
  properties: Property[]
}

export function PropertiesSection({ properties }: PropertiesSectionProps) {
  const [filtered, setFiltered] = useState<Property[]>(properties)

  function handleSearch(event: ChangeEvent<HTMLInputElement>) {
    const value = event.target.value

    if (value.trim() === "") {
      setFiltered(properties)
      return
    }

    const term = value.toLowerCase()
    setFiltered(properties.filter(property => property.address.mini_address.toLowerCase().includes(term)))
  }

  function handleFilterKind(value: string) {
    if (value === "Todos") {
      setFiltered(properties)
      return
    }

    setFiltered(properties.filter(property => property.kind === value))
  }

  function handleFilterStatus(value: string) {
    if (value === "Todos") {
      setFiltered(properties)
      return
    }

    setFiltered(properties.filter(property => property.status === value))
  }

  function handleClearFilter() {
    setFiltered(properties)
  }

  return (
    <section className="flex flex-col gap-6">
      <div className="flex items-center justify-between">
        <div className="w-full flex gap-2">
          <Input
            placeholder="Procurar..."
            onChange={handleSearch}
            className="w-4/5 sm:w-1/4"
          />
          <div className="hidden lg:flex gap-2">
            <Select onValueChange={handleFilterKind}>
              <SelectTrigger>
                <SelectValue placeholder="Tipo" defaultChecked />
              </SelectTrigger>
              <SelectContent className="text-sm md:text-base">
                <SelectItem value="Todos" defaultChecked>Todos</SelectItem>
                <SelectItem value="Residencial">Residencial</SelectItem>
                <SelectItem value="Comercial">Comercial</SelectItem>
                <SelectItem value="Industrial">Industrial</SelectItem>
                <SelectItem value="Terreno">Terreno</SelectItem>
                <SelectItem value="Rural">Rural</SelectItem>
              </SelectContent>
            </Select>

            <Select onValueChange={handleFilterStatus}>
              <SelectTrigger>
                <SelectValue placeholder="Status" defaultChecked />
              </SelectTrigger>
              <SelectContent className="text-sm md:text-base">
                <SelectItem value="Todos" defaultChecked>Todos</SelectItem>
                <SelectItem value="Disponível"><CircleCheck className="!size-4 text-emerald-500" />Disponível</SelectItem>
                <SelectItem value="Ocupado"><CircleMinus className="!size-4 text-yellow-500" />Ocupado</SelectItem>
                <SelectItem value="Indisponível"><CircleX className="!size-4 text-red-500" />Indisponível</SelectItem>
                <SelectItem value="Reservado"><LockKeyhole className="!size-4 text-sky-500" />Reservado</SelectItem>
                <SelectItem value="Reformando"><Hammer className="!size-4 text-orange-500" />Reformando</SelectItem>
              </SelectContent>
            </Select>

            <Button
              onClick={handleClearFilter}
              disabled={filtered.length === properties.length}
            >
              <CircleX className="!size-4" />
              Limpar Filtros
            </Button>
          </div>
        </div>

        <Link href="/dashboard/locacao/imoveis/novo">
          <Button>
            <Plus className="w-4 h-4" />
            <p className="hidden lg:block">Novo Imóvel</p>
          </Button>
        </Link>
      </div>
      <div className="grid gap-6 grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4">
        {filtered.map((property) => (
          <Card key={property.id} className="bg-zinc-900/20 backdrop-blur-2xl relative z-20 overflow-hidden">
            <div className={cn(["absolute z-10 w-10 h-10 blur-3xl bottom-0 right-0", bgColorStatusDetail(property.status)])} />
            <CardHeader className="flex flex-col gap-4 h-20">
              <CardTitle className="max-h-12">{property.address.mini_address}</CardTitle>
              <div className="flex flex-col gap-2 lg:flex-row lg:items-center">
                <StatusBadge status={property.status} />
                <Badge className="h-8" variant="outline">{property.kind}</Badge>
              </div>
            </CardHeader>
            <CardFooter className="flex justify-end">
              <div className="flex justify-end">
                <Link href={`/dashboard/locacao/imoveis/${property.id}`}>
                  <Button variant="outline">
                    <ArrowUpRight className="size-4" />
                    <p className="hidden sm:block">Detalhes</p>
                  </Button>
                </Link>
              </div>
            </CardFooter>
          </Card>
        ))}
      </div>
    </section>
  )
}