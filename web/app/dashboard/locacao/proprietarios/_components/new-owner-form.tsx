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

import { createOwner } from "@/actions/mutations/owner/create-owner"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { useRouter } from "next/navigation"
import { toast } from "sonner"
import { useHookFormMask } from "use-mask-input"

const owner_schema = z.object({
  fullname: z.string().min(10, "O nome deve conter ao menos 10 caracteres").max(100, "Máximo de 100 caracteres"),
  cpf: z.string().length(14),
  rg: z.string().min(5, "o rg deve ter no minimo 5 caracteres").max(15, "o rg deve ter no máximo 15 caracteres"),
  email: z.string().email("o e-mail digitado deve ser válido"),
  cellphone: z.string().min(14, "O telefone deve conter ao menos 10 dígitos").max(15, "O telefone deve conter no máximo 11 dígitos")
    .regex(/^\(\d{2}\)\s\d{4,5}-\d{4}$/, "O telefone deve conter apenas números"),
  occupation: z.string().min(3, "A ocupação deve conter ao menos 3 caracteres").max(50, "Máximo de 50 caracteres"),
  marital_status: z.enum(["Solteiro(a)", "Casado(a)", "Amasiado(a)", "Divorciado(a)", "Viúvo(a)", "União Estável"]),
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

export type OwnerRequest = z.infer<typeof owner_schema>

export function NewOwnerForm({ className, ...props }: React.ComponentProps<"form">) {
  const router = useRouter()

  const form = useForm<OwnerRequest>({
    resolver: zodResolver(owner_schema),
    defaultValues: {
      fullname: "",
      cpf: "",
      rg: "",
      email: "",
      cellphone: "",
      occupation: "",
      marital_status: "Solteiro(a)",
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

  async function onSubmit(values: OwnerRequest) {
    const status = await createOwner(values)

    switch (status) {
      case 201:
        toast.success("Proprietário criado com sucesso", { duration: 1500 })

        setTimeout(() => {
          router.push("/dashboard/locacao/proprietarios")
        }, 1500)
        break
      default:
        toast.error("Erro ao criar proprietário", { description: "Confira os dados e tente novamente", duration: 1500 })
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
          name="fullname"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Nome Completo</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Nome completo"
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
          name="cpf"
          render={({ field }) => (
            <FormItem>
              <FormLabel>CPF</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="CPF"
                  autoComplete="off"
                  {...field}
                  {...registerWithMask("cpf", '999.999.999-99', {
                    showMaskOnHover: false,
                    showMaskOnFocus: false,
                    required: true,
                  })}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="rg"
          render={({ field }) => (
            <FormItem>
              <FormLabel>RG</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="RG"
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
          name="cellphone"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Telefone / Celular</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Telefone ou celular"
                  autoComplete="off"
                  {...field}
                  {...registerWithMask("cellphone", ['(99) 9999-9999', '(99) 99999-9999'], {
                    showMaskOnHover: false,
                    showMaskOnFocus: false,
                    required: true,
                  })}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>E-mail</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="E-mail de contato"
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
          name="occupation"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Profissão / Ocupação</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Ocupação"
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
          name="marital_status"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Estado Civil</FormLabel>
              <FormControl>
                <Select
                  defaultValue={field.value}
                  onValueChange={field.onChange}
                  {...field}
                >
                  <SelectTrigger>
                    <SelectValue placeholder="Selecione o estado civil" />
                  </SelectTrigger>
                  <SelectContent className="text-sm md:text-base">
                    <SelectItem value="Solteiro(a)">Solteiro(a)</SelectItem>
                    <SelectItem value="Casado(a)">Casado(a)</SelectItem>
                    <SelectItem value="Amasiado(a)">Amasiado(a)</SelectItem>
                    <SelectItem value="Divorciado(a)">Divorciado(a)</SelectItem>
                    <SelectItem value="Viúvo(a)">Viúvo(a)</SelectItem>
                    <SelectItem value="União Estável">União Estável</SelectItem>
                  </SelectContent>
                </Select>
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <Separator className="col-span-1 sm:col-span-2" />

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
                  {...registerWithMask("address.zip_code", '99999999', {
                    showMaskOnHover: false,
                    showMaskOnFocus: false,
                    required: true,
                  })}
                />
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
