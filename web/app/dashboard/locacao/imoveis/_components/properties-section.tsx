'use client'

import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Card, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { cn } from "@/lib/utils"
import { Property } from "@/types/property"
import { ArrowUpRight, Plus } from "lucide-react"
import Link from "next/link"
import { ChangeEvent, useState } from "react"
import { colorStatusDetail, StatusBadge } from "./status-badge"

type PropertiesSectionProps = {
  properties: Property[]
}

export function PropertiesSection({ properties }: PropertiesSectionProps) {
  const [filtered, setFiltered] = useState<Property[]>(properties)

  function handleSearch(event: ChangeEvent<HTMLInputElement>) {
    if (event.target.value === "") {
      setFiltered(properties)
      return
    }

    const term = event.target.value.toLowerCase()
    const filter = filtered.filter(property => property.address.mini_address.toLowerCase().includes(term))
    setFiltered(filter)
  }

  function handleFilter(value: string) {
    if (value === "Todos") {
      setFiltered(properties)
      return
    }

    setFiltered(properties.filter(property => property.kind === value))
  }

  return (
    <section className="flex flex-col gap-4">
      <div className="flex items-center justify-between">
        <div className="w-full flex gap-2">
          <Input
            placeholder="Procurar..."
            onChange={handleSearch}
            className="w-4/5 sm:w-1/4"
          />
          <div className="hidden lg:flex">
            <Select onValueChange={handleFilter}>
              <SelectTrigger>
                <SelectValue placeholder="Todos" defaultChecked />
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
          </div>
        </div>

        <Link href="/dashboard/locacao/imoveis/novo">
          <Button>
            <Plus className="w-4 h-4" />
            <p className="hidden lg:block">Novo Imóvel</p>
          </Button>
        </Link>
      </div>
      <div className="grid gap-6 grid-cols-1 sm:grid-cols-2 xl:grid-cols-3">
        {filtered.map((property) => (
          <Card key={property.id} className="bg-zinc-900/20 backdrop-blur-2xl relative z-20 overflow-hidden">
            <div className={cn(["absolute z-10 w-10 h-10 blur-3xl bottom-0 right-0", `bg-${colorStatusDetail(property.status)}`])} />
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