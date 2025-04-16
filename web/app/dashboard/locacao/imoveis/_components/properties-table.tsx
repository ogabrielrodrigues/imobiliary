'use client'

import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Property } from "@/types/property"
import { ColumnDef } from "@tanstack/react-table"
import { ArrowUpRight, Check, CircleCheck, CircleX } from "lucide-react"
import { DataTable } from "./properties-data-table"

interface PropertiesTableProps {
  properties: Property[]
}

const columns: ColumnDef<Property>[] = [
  {
    header: "Endereço",
    accessorKey: "address.mini_address",
    id: "address",
  },
  {
    header: "Status",
    accessorKey: "status",
    cell: ({ row }) => {
      const status = row.original.status.toUpperCase()

      switch (status) {
        case "DISPONÍVEL":
          return <Badge variant="outline" className="h-8">
            <CircleCheck className="!size-4 text-green-500" />{status}
          </Badge >
        case "OCCUPIED":
          return <Badge variant="outline" className="h-8">
            <Check className="!size-4 text-yellow-500" />{status}
          </Badge >
        case "UNAVAILABLE":
          return <Badge variant="outline" className="h-8">
            <CircleX className="!size-4 text-red-500" />{status}
          </Badge >
      }
    },
  },
  {
    header: "Tipo",
    accessorKey: "address.kind",
  },
  {
    header: "Cód. Água",
    accessorKey: "water_id",
  },
  {
    header: "Cód. Energia",
    accessorKey: "energy_id",
  },
  {
    id: "actions",
    cell: () => (
      <Button variant="outline">
        <ArrowUpRight className="size-4" />
        <p className="text-xs hidden md:block">Ver</p>
      </Button>
    ),
  },
]

export function PropertiesTable({ properties }: PropertiesTableProps) {
  return (
    <DataTable
      className="mt-6"
      columns={columns}
      data={properties}
    />
  )
}
