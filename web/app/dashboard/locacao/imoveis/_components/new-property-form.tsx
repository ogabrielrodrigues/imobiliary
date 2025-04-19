'use client'

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"

import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"

import { cn } from "@/lib/utils"

import { Separator } from "@/components/ui/separator"

import { createProperty } from "@/actions/properties"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { useRouter } from "next/navigation"
import { toast } from "sonner"

export const property_schema = z.object({
  water_id: z.string().min(3, "O código da água deve conter ao menos 3 caracteres"),
  energy_id: z.string().min(3, "O código da energia deve conter ao menos 3 caracteres"),
  status: z.enum(['Disponível', 'Ocupado', 'Indisponível', 'Reservado', 'Reformando']),
  kind: z.enum(['Residencial', 'Comercial', 'Industrial', 'Terreno', 'Rural']),
  address: z.object({
    street: z.string().min(3, "O nome da rua deve conter ao menos 3 caracteres"),
    number: z.string().min(1, "O número da rua deve conter ao menos 1 caracter"),
    neighborhood: z.string().min(3, "O bairro deve conter ao menos 3 caracteres"),
    complement: z.string().optional(),
    city: z.string().min(3, "A cidade deve conter ao menos 3 caracteres"),
    state: z.string().optional(),
    zip_code: z.string().min(8, "O CEP deve conter ao menos 8 caracteres"),
  }),
})

export function NewPropertyForm({ className, ...props }: React.ComponentProps<"form">) {
  const router = useRouter()

  const form = useForm<z.infer<typeof property_schema>>({
    resolver: zodResolver(property_schema),
    defaultValues: {
      water_id: "",
      energy_id: "",
      status: "Disponível",
      kind: "Residencial",
      address: {
        street: "",
        number: "",
        neighborhood: "",
        complement: "",
        city: "",
        state: "",
        zip_code: ""
      }
    }
  })

  async function onSubmit(values: z.infer<typeof property_schema>) {
    const status = await createProperty(values)

    switch (status) {
      case 201:
        toast.success("Imóvel criado com sucesso", { duration: 1500 })

        setTimeout(() => {
          router.push("/dashboard/locacao/imoveis")
        }, 1500)
        break
      default:
        toast.error("Erro ao criar imóvel", { description: "Confira os dados e tente novamente", duration: 1500 })
        break
    }
  }

  return (
    <Form {...form}>
      <form
        className={cn("grid grid-cols-subgrid md:!grid-cols-2 space-y-4 gap-x-2", className)}
        {...props}
        onSubmit={form.handleSubmit(onSubmit)}
      >
        <FormField
          control={form.control}
          name="address.street"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Logradouro</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Logradouro"
                  autoComplete="off"
                  autoFocus
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="address.number"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Número</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Número"
                  autoComplete="off"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="address.complement"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Complemento</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Complemento"
                  autoComplete="off"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="address.neighborhood"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Bairro</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Bairro"
                  autoComplete="off"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="address.city"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Cidade</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Cidade"
                  autoComplete="off"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="address.zip_code"
          render={({ field }) => (
            <FormItem>
              <FormLabel>CEP</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="CEP"
                  autoComplete="off"
                  maxLength={8}
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <Separator className="col-span-1 sm:col-span-2" />

        <FormField
          control={form.control}
          name="water_id"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Cód. Água</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Cód. água"
                  autoComplete="off"
                  autoFocus
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="energy_id"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Cód. Energia</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Cód. energia"
                  autoComplete="off"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <Separator className="col-span-1 sm:col-span-2" />

        <FormField
          control={form.control}
          name="kind"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Tipo</FormLabel>
              <FormControl>
                <Select
                  defaultValue={field.value}
                  onValueChange={field.onChange}
                  {...field}
                >
                  <SelectTrigger>
                    <SelectValue placeholder="Selecione o tipo" />
                  </SelectTrigger>
                  <SelectContent className="text-sm md:text-base">
                    <SelectItem value="Residencial">Residencial</SelectItem>
                    <SelectItem value="Comercial">Comercial</SelectItem>
                    <SelectItem value="Industrial">Industrial</SelectItem>
                    <SelectItem value="Terreno">Terreno</SelectItem>
                    <SelectItem value="Rural">Rural</SelectItem>
                  </SelectContent>
                </Select>
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="status"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Status</FormLabel>
              <FormControl>
                <Select
                  defaultValue={field.value}
                  onValueChange={field.onChange}
                  {...field}
                >
                  <SelectTrigger>
                    <SelectValue placeholder="Selecione o tipo" />
                  </SelectTrigger>
                  <SelectContent className="text-sm md:text-base">
                    <SelectItem value="Disponível">Disponível</SelectItem>
                    <SelectItem value="Ocupado">Ocupado</SelectItem>
                    <SelectItem value="Indisponível">Indisponível</SelectItem>
                    <SelectItem value="Reservado">Reservado</SelectItem>
                    <SelectItem value="Reformando">Reformando</SelectItem>
                  </SelectContent>
                </Select>
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit" className="sm:col-start-2">
          Cadastrar
        </Button>
      </form>
    </Form>
  )
}
