'use client'

import { sendMail } from "@/actions/mutation/send-mail";
import { Button } from "@/components/ui/button";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from 'zod';

const contact_schema = z.object({
  name: z.string().min(10, "Nome é obrigatório"),
  email: z.string().email("E-mail inválido"),
  message: z.string().min(5, "Mensagem é obrigatória"),
})

export type ContactFormValues = z.infer<typeof contact_schema>

export function ContactForm() {
  const form = useForm({
    resolver: zodResolver(contact_schema),
    defaultValues: {
      name: "",
      email: "",
      message: "",
    }
  })

  async function onSubmit(values: ContactFormValues) {
    const status = await sendMail(values)

    switch (status) {
      case 250:
        form.reset()
        toast.success("Mensagem enviada com sucesso!")
        break
      case 500:
        toast.error("Erro ao enviar mensagem. Tente novamente mais tarde.")
        break
    }
  }

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="w-full flex flex-col gap-4"
      >
        <div className="grid gap-4 sm:grid-cols-2">
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Nome</FormLabel>
                <FormControl>
                  <Input
                    type="text"
                    placeholder="Seu nome completo"
                    {...field}
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
                    type="email"
                    placeholder="Seu e-mail cadastrado"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )} />
        </div>
        <FormField
          control={form.control}
          name="message"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Mensagem</FormLabel>
              <FormControl>
                <Textarea
                  placeholder="Descreva o motivo de seu contato"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )} />

        <Button type="submit" className="font-bold">
          Enviar
        </Button>
      </form>
    </Form>
  );
}