import { Check } from "lucide-react";

export function FreePlan() {
  return (
    <ul className="space-y-1">
      <li className="flex items-center gap-2 text-sm">
        <Check className="size-4 text-emerald-500" />
        Gerenciamento de até 30 imóveis
      </li>
      <li className="flex items-center gap-2 text-sm">
        <Check className="size-4 text-emerald-500" />
        Suporte via e-mail
      </li>
    </ul>
  )
}

export function ProPlan() {
  return (
    <ul className="space-y-1">
      <li className="flex items-center gap-2 text-sm">
        <Check className="size-4 text-emerald-500" />
        Gerenciamento ilimitado de imóveis
      </li>
      <li className="flex items-center gap-2 text-sm">
        <Check className="size-4 text-emerald-500" />
        Notificações por e-mail
      </li>
      <li className="flex items-center gap-2 text-sm">
        <Check className="size-4 text-emerald-500" />
        Levantamento de taxas de consumo
      </li>
      <li className="flex items-center gap-2 text-sm">
        <Check className="size-4 text-emerald-500" />
        Gerencimento de vistorias
      </li>
      <li className="flex items-center gap-2 text-sm">
        <Check className="size-4 text-emerald-500" />
        Administração de carteira
      </li>
      <li className="flex items-center gap-2 text-sm">
        <Check className="size-4 text-emerald-500" />
        Métricas e relatórios detalhados
      </li>
      <li className="flex items-center gap-2 text-sm">
        <Check className="size-4 text-emerald-500" />
        Até 3 usuários
      </li>
      <li className="flex items-center gap-2 text-sm">
        <Check className="size-4 text-emerald-500" />
        Suporte via e-mail e chat
      </li>
    </ul>
  )
}


