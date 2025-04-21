import { Badge } from "@/components/ui/badge"
import { cn } from "@/lib/utils"
import { Property } from "@/types/property"
import { CircleCheck, CircleMinus, CircleX, Hammer, LockKeyhole } from "lucide-react"

type StatusBadgeProps = Pick<Property, 'status'>

export function bgColorStatusDetail(status: string) {
  switch (status) {
    case "Disponível":
      return "bg-emerald-500"
    case "Ocupado":
      return "bg-yellow-500"
    case "Indisponível":
      return "bg-red-500"
    case "Reformando":
      return "bg-orange-500"
    case "Reservado":
      return "bg-sky-500"
  }
}

export function textColorStatusDetail(status: string) {
  switch (status) {
    case "Disponível":
      return "text-emerald-500"
    case "Ocupado":
      return "text-yellow-500"
    case "Indisponível":
      return "text-red-500"
    case "Reformando":
      return "text-orange-500"
    case "Reservado":
      return "text-sky-500"
  }
}

export function StatusBadge({ status }: StatusBadgeProps) {
  switch (status) {
    case "Disponível":
      return <Badge variant="outline" className="h-8 gap-2">
        <CircleCheck className={cn(["!size-4", textColorStatusDetail(status)])} />{status}
      </Badge >
    case "Ocupado":
      return <Badge variant="outline" className="h-8 gap-2">
        <CircleMinus className={cn(["!size-4", textColorStatusDetail(status)])} />{status}
      </Badge >
    case "Indisponível":
      return <Badge variant="outline" className="h-8 gap-2">
        <CircleX className={cn(["!size-4", textColorStatusDetail(status)])} />{status}
      </Badge >
    case "Reservado":
      return <Badge variant="outline" className="h-8 gap-2">
        <LockKeyhole className={cn(["!size-4", textColorStatusDetail(status)])} />{status}
      </Badge >
    case "Reformando":
      return <Badge variant="outline" className="h-8 gap-2">
        <Hammer className={cn(["!size-4", textColorStatusDetail(status)])} />{status}
      </Badge >
  }
}