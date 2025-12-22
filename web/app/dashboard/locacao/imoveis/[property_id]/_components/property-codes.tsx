'use client'

import { Button } from "@/components/ui/button"
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip"
import { Property } from "@/types/property"
import { ClipboardCheck } from "lucide-react"
import { toast } from "sonner"

type PropertyCodesProps = {
  property: Property
}

export function PropertyCodes({ property }: PropertyCodesProps) {
  function copyCodeToClipboard(code: string) {
    navigator.clipboard.writeText(code)

    toast("Código copiado para área de transferência", {
      duration: 3000,
      icon: <ClipboardCheck />,
      description: "Use o código na respectiva concessionária para consultar taxas de consumo e débitos."
    })
  }

  return (
    <div className="flex flex-col sm:flex-row gap-2 sm:gap-6 text-xs text-muted-foreground">
      <span>Cód. Água - <Tooltip>
        <TooltipTrigger asChild>
          <Button onClick={() => copyCodeToClipboard(property.water_id)} variant="link" className="!p-0 !text-xs">{property.water_id}</Button>
        </TooltipTrigger>
        <TooltipContent side="bottom" className="text-muted-foreground">Copiar p/ área de transferência</TooltipContent>
      </Tooltip></span>
      <span>Cód. Energia - <Tooltip>
        <TooltipTrigger asChild>
          <Button onClick={() => copyCodeToClipboard(property.energy_id)} variant="link" className="!p-0 !text-xs">{property.energy_id}</Button>
        </TooltipTrigger>
        <TooltipContent side="bottom" className="text-muted-foreground">Copiar p/ área de transferência</TooltipContent>
      </Tooltip></span>
    </div>
  )
}