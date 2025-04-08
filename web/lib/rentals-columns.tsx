import { DragHandle } from "@/components/drag-handle";
import { Checkbox } from "@radix-ui/react-checkbox";
import { ColumnDef } from "@tanstack/react-table";
import { Badge } from "@/components/ui/badge";
import { z } from "zod";
import { ArrowUpDown, CircleCheck, Eye, Loader, TriangleAlert } from "lucide-react";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";

export const Rentals = z.object({
  id: z.number(),
  nomecompleto: z.string(),
  valor: z.number(),
  status: z.string(),
  vencimento: z.string(),
})

export const rentals_columns: ColumnDef<z.infer<typeof Rentals>>[] = [
  {
    id: "drag",
    header: () => null,
    cell: ({ row }) => <DragHandle id={row.original.id} />,
  },
  {
    id: "select",
    header: ({ table }) => (
      <div className="flex items-center justify-center">
        <Checkbox
          checked={
            table.getIsAllPageRowsSelected() ||
            (table.getIsSomePageRowsSelected() && "indeterminate")
          }
          onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
          aria-label="Selecionar tudo"
        />
      </div>
    ),
    cell: ({ row }) => (
      <div className="flex items-center justify-center">
        <Checkbox
          checked={row.getIsSelected()}
          onCheckedChange={(value) => row.toggleSelected(!!value)}
          aria-label="Selecionar linha"
        />
      </div>
    ),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "nomecompleto",
    header: "Nome Completo",
    cell: ({ row }) => (
      <span>{row.original.nomecompleto}</span>
    ),
    enableHiding: false,
  },
  {
    accessorKey: "valor",
    header: "Valor",
    cell: ({ row }) => (
      <div className="w-32">
        R${row.original.valor.toString()}
      </div>
    ),
  },
  {
    accessorKey: "status",
    header: "Status atual",
    cell: ({ row }) => (
      <Badge variant="outline" className="text-muted-foreground py-2 text-sm">
        {row.original.status === "Pago" && <CircleCheck className="!size-4 stroke-emerald-500" />}
        {row.original.status === "A vencer" && <Loader className="!size-4 stroke-yellow-500" />}
        {row.original.status === "Atrasado" && <TriangleAlert className="!size-4 stroke-red-500" />}
        {row.original.status}
      </Badge>
    ),
  },
  {
    accessorKey: "vencimento",
    header: ({ column }) => (
      <Button
        variant="ghost"
        onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
      >
        Vencimento
        <ArrowUpDown className="ml-2 h-4 w-4" />
      </Button>
    ),
    cell: ({ row }) => (
      <div>{row.original.vencimento}</div>
    ),
  },
  {
    id: "actions",
    header: "Ação",
    cell: ({ row }) => (
      <Link href={`/dashboard/imoveis/alugueis/${row.original.id}`}>
        <Tooltip>
          <TooltipTrigger asChild>
            <Button variant="outline" size="icon" className="text-muted-foreground cursor-pointer">
              <Eye />
            </Button>
          </TooltipTrigger>
          <TooltipContent side="bottom">
            Visualizar
          </TooltipContent>
        </Tooltip>
      </Link>
    ),
  },
]