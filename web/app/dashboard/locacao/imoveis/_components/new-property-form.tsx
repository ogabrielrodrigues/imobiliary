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

import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
export const property_schema = z.object({
  water_id: z.string().min(3, "O código da água deve conter ao menos 3 caracteres"),
  energy_id: z.string().min(3, "O código da energia deve conter ao menos 3 caracteres"),
  address: z.object({
    street: z.string().min(3, "O nome da rua deve conter ao menos 3 caracteres"),
    number: z.string().min(1, "O número da rua deve conter ao menos 1 caracter"),
    neighborhood: z.string().min(3, "O bairro deve conter ao menos 3 caracteres"),
    complement: z.string().optional(),
    city: z.string().min(3, "A cidade deve conter ao menos 3 caracteres"),
    state: z.string().min(2, "O estado deve conter ao menos 2 caracteres"),
    state_abbr: z.string().min(2, "A sigla do estado deve conter ao menos 2 caracteres"),
    zip_code: z.string().min(8, "O CEP deve conter ao menos 8 caracteres"),
    kind: z.enum(["Residencial", "Comercial"]),
  }),
})

export function NewPropertyForm({ className, ...props }: React.ComponentProps<"form">) {
  const form = useForm<z.infer<typeof property_schema>>({
    resolver: zodResolver(property_schema),
    defaultValues: {
      water_id: "",
      energy_id: "",
      address: {
        street: "",
        number: "",
        neighborhood: "",
        city: "",
        state: "",
        state_abbr: "",
        zip_code: "",
        kind: "Residencial"
      }
    }
  })

  async function onSubmit(values: z.infer<typeof property_schema>) {
    console.log(values)
  }

  return (
    <Form {...form}>
      <form
        className={cn("grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 space-y-4 gap-x-2", className)}
        {...props}
        onSubmit={form.handleSubmit(onSubmit)}
      >
        <FormField
          control={form.control}
          name="address.street"
          render={({ field }) => (
            <FormItem className="col-span-1 xl:col-span-2">
              <FormLabel>Rua</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Rua"
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
          name="address.number"
          render={({ field }) => (
            <FormItem className="col-span-1">
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
          name="address.neighborhood"
          render={({ field }) => (
            <FormItem className="col-span-1">
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
            <FormItem className="col-span-1">
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
          name="address.state"
          render={({ field }) => (
            <FormItem className="col-span-1">
              <FormLabel>Estado</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Estado"
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
          name="address.state_abbr"
          render={({ field }) => (
            <FormItem className="col-span-1">
              <FormLabel>Sigla do Estado</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Sigla"
                  autoComplete="off"
                  maxLength={2}
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
            <FormItem className="col-span-1">
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

        <Separator className="col-span-1 sm:col-span-2 xl:col-span-3" />

        <FormField
          control={form.control}
          name="water_id"
          render={({ field }) => (
            <FormItem className="col-span-1">
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
            <FormItem className="col-span-1">
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
        <FormField
          control={form.control}
          name="address.kind"
          render={({ field }) => (
            <FormItem className="col-span-1">
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
                  </SelectContent>
                </Select>
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit" className="w-full col-start-2 col-end-2 xl:col-start-3 xl:col-end-3">
          Cadastrar
        </Button>
      </form>
    </Form>
  )
}
