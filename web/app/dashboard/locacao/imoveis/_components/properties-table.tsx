'use client'

import { Button } from "@/components/ui/button"
import { Property } from "@/types/property"
import { ColumnDef } from "@tanstack/react-table"
import { ArrowUpDown, ArrowUpRight } from "lucide-react"
import Link from "next/link"
import { DataTable } from "./properties-data-table"
import { StatusBadge } from "./status-badge"

interface PropertiesTableProps {
  properties: Property[]
}

const columns: ColumnDef<Property>[] = [
  {
    header: ({ column }) => {
      return (
        <Button
          variant="ghost"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Endereço
          <ArrowUpDown className="ml-2 h-4 w-4" />
        </Button>
      )
    },
    accessorKey: "address.mini_address",
    id: "address",
  },
  {
    header: "Status",
    accessorKey: "status",
    cell: ({ row }) => <StatusBadge status={row.original.status} />,
  },
  {
    header: "Tipo",
    accessorKey: "kind",
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
    cell: ({ row }) => (
      <Link href={`/dashboard/locacao/imoveis/${row.original.id}`}>
        <Button variant="outline">
          <ArrowUpRight className="size-4" />
          <p className="text-xs hidden md:block">Ver</p>
        </Button>
      </Link>
    ),
  },
]

export function PropertiesTable({ properties }: PropertiesTableProps) {
  return (
    <DataTable
      columns={columns}
      data={properties}
    />
  )
}
