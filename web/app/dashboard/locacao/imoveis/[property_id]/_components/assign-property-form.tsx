'use client'

import { assignOwner } from "@/actions/mutations/owner/assign-owner";
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Owner } from "@/types/owner";
import { zodResolver } from "@hookform/resolvers/zod";
import Link from "next/link";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";

type AssignPropertyFormProps = {
  owners: Owner[] | undefined
  property_id: string
}

const assign_property_schema = z.object({
  owner_id: z.string().uuid()
})

export type AssignOwnerRequestData = z.infer<typeof assign_property_schema>

export function AssignPropertyForm({ property_id, owners }: AssignPropertyFormProps) {
  const form = useForm<AssignOwnerRequestData>({
    resolver: zodResolver(assign_property_schema),
    defaultValues: {
      owner_id: "",
    }
  })

  async function onSubmit(values: AssignOwnerRequestData) {
    const status = await assignOwner({
      id: property_id,
      data: values
    })

    switch (status) {
      case 200:
        toast.success("Proprietário associado com sucesso", { duration: 1500 })
        setTimeout(() => {
          window.location.reload()
        }, 1500)
        break
      default:
        toast.error("Erro ao associar proprietário", { description: "Confira os dados e tente novamente" })
        break
    }
  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>
          Associar proprietário
        </Button>
      </DialogTrigger>
      <DialogContent className="space-y-6">
        <DialogTitle className="text-2xl font-bold">Associar Proprietário</DialogTitle>
        {!owners || owners.length === 0 ? (
          <div className="w-full text-sm text-muted-foreground flex items-center flex-col gap-4">
            Nenhum proprietário cadastrado<br className="xl:hidden" />
            <Link href="/dashboard/locacao/proprietarios/novo" className="text-primary">
              <Button>
                Criar novo
              </Button>
            </Link>
          </div>
        ) : (
          <Form {...form}>
            <form
              className="m-0 gap-8 flex flex-col flex-1"
              onSubmit={form.handleSubmit(onSubmit)}
            >
              <FormField
                control={form.control}
                name="owner_id"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Proprietário</FormLabel>
                    <FormControl>
                      <Select
                        {...field}
                        defaultValue={field.value}
                        onValueChange={field.onChange}
                      >
                        <SelectTrigger>
                          <SelectValue placeholder="Selecione um proprietário" />
                        </SelectTrigger>
                        <SelectContent>
                          {owners.map((owner) => (
                            <SelectItem key={owner.id} value={owner.id}>{owner.fullname}</SelectItem>
                          ))}
                        </SelectContent>
                      </Select>
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <div className="w-full flex justify-between items-center">
                <p className="text-sm text-muted-foreground flex flex-col lg:flex-row lg:gap-2">
                  Não achou o proprietário?<br className="xl:hidden" />
                  <Link href="/dashboard/locacao/proprietarios/novo" className="text-primary">
                    Criar novo
                  </Link>
                </p>

                <Button type="submit">
                  Associar
                </Button>
              </div>
            </form>
          </Form>
        )}
      </DialogContent>
    </Dialog>
  )
}