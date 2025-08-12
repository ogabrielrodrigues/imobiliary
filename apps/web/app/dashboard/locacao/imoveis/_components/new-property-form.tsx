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

import { createProperty } from "@/actions/mutations/property/create-property"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { searchCEP } from "@/lib/cep/api-brasil"
import { Owner } from "@/types/owner"
import { LoaderCircle } from "lucide-react"
import { useRouter } from "next/navigation"
import { useState } from "react"
import { toast } from "sonner"
import { useHookFormMask } from "use-mask-input"

export const property_schema = z.object({
  owner_id: z.string().uuid("Você deve selecionar um proprietário"),
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

export type CreatePropertyRequest = z.infer<typeof property_schema>

export type NewPropertyFormProps = React.ComponentProps<"form"> & {
  owners: Owner[]
}

export function NewPropertyForm({ owners, className, ...props }: NewPropertyFormProps) {
  const router = useRouter()
  const [loading, setLoading] = useState(false)

  const form = useForm<CreatePropertyRequest>({
    resolver: zodResolver(property_schema),
    defaultValues: {
      owner_id: "",
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

  const registerWithMask = useHookFormMask(form.register);

  async function searchAddressByCEP(cep: string) {
    const address = await searchCEP(cep)

    if (address) {
      form.setValue("address.street", address.street)
      form.setValue("address.neighborhood", address.neighborhood)
      form.setValue("address.city", address.city)
      form.setValue("address.state", address.state)
    } else {
      toast.error("CEP não encontrado", { duration: 1500 })
    }
  }

  async function onSubmit(values: CreatePropertyRequest) {
    setLoading(true)

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

    setLoading(false)
  }

  return (
    <Form {...form}>
      <form
        className="flex flex-col pb-4 gap-4"
        {...props}
        onSubmit={form.handleSubmit(onSubmit)}
      >
        <div className={cn("grid grid-cols-subgrid md:!grid-cols-2 gap-y-4 gap-x-2", className)}>
          <FormField
            control={form.control}
            name="owner_id"
            render={({ field }) => (
              <FormItem className="col-span-2">
                <FormLabel>Proprietário</FormLabel>
                <FormControl>
                  <Select
                    defaultValue={field.value}
                    onValueChange={field.onChange}
                    disabled={loading || owners.length === 0}
                    {...field}
                  >
                    <SelectTrigger autoFocus>
                      <SelectValue
                        placeholder={owners.length > 0
                          ? "Selecione o proprietário"
                          : "Sem proprietários"}
                      />
                    </SelectTrigger>
                    <SelectContent className="text-sm md:text-base">
                      {owners.length > 0 && owners.map(owner => <SelectItem key={owner.id} value={owner.id}>{owner.fullname}</SelectItem>)}
                    </SelectContent>
                  </Select>
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
                    autoFocus
                    disabled={loading}
                    maxLength={8}
                    {...field}
                    {...registerWithMask("address.zip_code", '99999999', {
                      showMaskOnHover: false,
                      showMaskOnFocus: false,
                      required: true,
                      onBlur: (e) => {
                        searchAddressByCEP(e.target.value)
                      }
                    })}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

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
                    disabled={loading}
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
                    disabled={loading}
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
                    disabled={loading}
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
                    disabled={loading}
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
                    disabled={loading}
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

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
                    disabled={loading}
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
                    disabled={loading}
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

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
                    disabled={loading}
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
                    disabled={loading}
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
        </div>

        <div className="flex justify-end">
          <Button disabled={loading} type="submit" size="lg" className="">
            {loading && <LoaderCircle className="size-4 animate-spin" />}Cadastrar
          </Button>
        </div>
      </form>
    </Form>
  )
}
