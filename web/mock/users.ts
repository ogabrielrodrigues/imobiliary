import { User } from "@/types/user";

export const users: User[] = [
  {
    id: "550e8400-e29b-41d4-a716-446655440000",
    fullname: "João da Silva",
    creci: "67543-F",
    telefone: "11 98765-4321",
    email: "joao.silva@imobdesk.com",
    password: "password",
    avatar: "https://github.com/ogabrielrodrigues.png",
    plan: {
      kind: "free",
      propertiesTotalQuota: 30,
      propertiesUsedQuota: 5,
      propertiesRemainingQuota: 25,
    },
  },
];
