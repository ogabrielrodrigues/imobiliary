'use client'
import { Button } from "./ui/button";

import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { User } from "@/types/user";
import { Check, Sparkles } from "lucide-react";
import { PropsWithChildren } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";

type ProPlanDialogProps = PropsWithChildren & {
  user: User | undefined
}

export function ProPlanDialog({ user, children }: ProPlanDialogProps) {
  return (
    <Dialog>
      <DialogTrigger asChild>
        {children}
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Atualizar para o PRO</DialogTitle>
          <DialogDescription>
            Mude para o plano <strong className="text-primary">PRO</strong> e tenha acesso a todas as funcionalidades do Imobiliary.
          </DialogDescription>
        </DialogHeader>
        <Card>
          <CardHeader>
            <CardTitle>Plano PRO</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <ul className="space-y-1">
              <li className="flex items-center gap-2 text-sm">
                <Check className="size-4 text-emerald-500" />
                Gerenciamento ilimitado de imóveis
              </li>
              <li className="flex items-center gap-2 text-sm">
                <Check className="size-4 text-emerald-500" />
                Notificações por e-mail e push
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
                Até 5 usuários
              </li>
              <li className="flex items-center gap-2 text-sm">
                <Check className="size-4 text-emerald-500" />
                Suporte 24/7 via e-mail e chat
              </li>
            </ul>
          </CardContent>
        </Card>
        <DialogFooter className="!flex-row !justify-between">
          <div className="flex items-center">
            <h1 className="text-xl sm:text-2xl font-bold">R$15,99</h1>
            <p className="text-xl text-muted-foreground">/mês</p>
          </div>

          <Button className="w-1/2">
            <Sparkles />
            Atualizar Agora
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}