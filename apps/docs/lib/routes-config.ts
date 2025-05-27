// for page navigation & to sort on leftbar

import { TagVariant } from "@/components/sublink";

export type EachRoute = {
  title: string;
  href: string;
  noLink?: true; // noLink will create a route segment (section) but cannot be navigated
  items?: EachRoute[];
  tag?: {
    title: string;
    variant: TagVariant;
  }
};

export const ROUTES: EachRoute[] = [
  {
    title: "API",
    href: "/api",
    items: [
      { title: "Autenticação", href: "/autenticacao" },
      { title: "Respostas", href: "/respostas" },
      {
        title: "Rotas",
        href: "/rotas",
        items: [
          {
            title: "Usuário",
            href: "/usuario",
            items: [
              { title: "Listar Usuários", href: "/listar", tag: { title: "GET", variant: "get" } },
              { title: "Buscar Usuário", href: "/buscar", tag: { title: "GET", variant: "get" } },
              { title: "Criar Usuário", href: "/criar", tag: { title: "POST", variant: "post" } },
              { title: "Autenticar Usuário", href: "/autenticar", tag: { title: "POST", variant: "post" } },
              { title: "Alterar Avatar", href: "/alterar-avatar", tag: { title: "PUT", variant: "put" } },
            ],
          },
          {
            title: "Imóvel",
            href: "/imovel",
            items: [
              { title: "Listar Imóveis", href: "/listar", tag: { title: "GET", variant: "get" } },
              { title: "Buscar Imóvel", href: "/buscar", tag: { title: "GET", variant: "get" } },
              { title: "Criar Imóvel", href: "/criar", tag: { title: "POST", variant: "post" } },
            ],
          },
          {
            title: "Proprietário",
            href: "/proprietario",
            items: [
              { title: "Listar Proprietários", href: "/listar", tag: { title: "GET", variant: "get" } },
              { title: "Buscar Proprietário", href: "/buscar", tag: { title: "GET", variant: "get" } },
              { title: "Criar Proprietário", href: "/criar", tag: { title: "POST", variant: "post" } },
              { title: "Atribuir Proprietário", href: "/atribuir", tag: { title: "PUT", variant: "put" } },
            ],
          },
        ]
      },
    ],
  },
];

type Page = { title: string; href: string };

function getRecurrsiveAllLinks(node: EachRoute) {
  const ans: Page[] = [];
  if (!node.noLink) {
    ans.push({ title: node.title, href: node.href });
  }
  node.items?.forEach((subNode) => {
    const temp = { ...subNode, href: `${node.href}${subNode.href}` };
    ans.push(...getRecurrsiveAllLinks(temp));
  });
  return ans;
}

export const page_routes = ROUTES.map((it) => getRecurrsiveAllLinks(it)).flat();
