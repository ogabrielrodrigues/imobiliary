'use client'
import { Button } from "./ui/button";

import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { Sparkles } from "lucide-react";
import { PropsWithChildren } from "react";
import { ProPlan } from "./plan-describe";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";

export function ProPlanDialog({ children }: PropsWithChildren) {
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
            <ProPlan />
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