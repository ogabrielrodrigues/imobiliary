import { capitalizeFirstLetter } from "./utils";

type DynamicParam = {
  name: string;
  isDynamic: boolean;
  children?: RouteStructure;
}

type RouteNode = {
  name: string;
  href: string;
  children?: RouteStructure;
  dynamicChildren?: Record<string, DynamicParam>;
}

type RouteStructure = {
  [key: string]: RouteNode;
}

type BreadcrumbItem = {
  name: string;
  href: string;
  isLast: boolean;
  dynamicId?: string;
}

function isDynamicSegment(segment: string): boolean {
  return (
    /^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i.test(segment)
  );
};

type GenerateBreadcrumbPathParams = {
  pathname: string,
  routeStructure: RouteStructure
}

export function generateBreadcrumbPath({ pathname, routeStructure }: GenerateBreadcrumbPathParams): BreadcrumbItem[] {
  if (!pathname) return [];

  const pathSegments = pathname.split('/').filter(Boolean);
  if (pathSegments.length === 0) return [];

  const breadcrumbItems: BreadcrumbItem[] = [];

  const dashboardNode: RouteNode = routeStructure.dashboard;

  if (!dashboardNode) {
    return [];
  }

  let currentPath = '';
  let currentStructure: RouteStructure | undefined = routeStructure;
  let dynamicParent: RouteNode | null = null;

  for (let i = 0; i < pathSegments.length; i++) {
    const segment = pathSegments[i];
    currentPath += `/${segment}`;

    if (i === 0 && segment === 'dashboard') {
      breadcrumbItems.push({
        name: dashboardNode.name,
        href: dashboardNode.href,
        isLast: pathSegments.length === 1
      });

      currentStructure = dashboardNode.children;
    } else if (currentStructure && segment in currentStructure) {
      const currentNode: RouteNode = currentStructure[segment];

      breadcrumbItems.push({
        name: currentNode.name,
        href: currentNode.href,
        isLast: i === pathSegments.length - 1
      });

      dynamicParent = currentNode;
      currentStructure = currentNode.children;
    } else if (
      isDynamicSegment(segment) &&
      dynamicParent &&
      dynamicParent.dynamicChildren &&
      'id' in dynamicParent.dynamicChildren
    ) {
      const dynamicConfig: DynamicParam = dynamicParent.dynamicChildren.id;

      if (dynamicConfig) {
        breadcrumbItems.push({
          name: dynamicConfig.name,
          href: currentPath,
          isLast: i === pathSegments.length - 1,
          dynamicId: segment
        });

        currentStructure = dynamicConfig.children;
      }
    } else {
      breadcrumbItems.push({
        name: capitalizeFirstLetter(segment),
        href: currentPath,
        isLast: i === pathSegments.length - 1
      });

      currentStructure = undefined;
    }
  }

  return breadcrumbItems;
};

export const routes: RouteStructure = {
  dashboard: {
    name: "Visão Geral",
    href: "/dashboard",
    children: {
      conta: {
        name: "Conta",
        href: "/dashboard/conta",
      },
      locacao: {
        name: "Alugueis",
        href: "/dashboard/locacao",
        children: {
          inquilinos: {
            name: "Inquilinos",
            href: "/dashboard/locacao/inquilinos",
            dynamicChildren: {
              id: {
                name: "Perfil do Inquilino",
                isDynamic: true
              }
            }
          },
          imoveis: {
            name: "Imóveis",
            href: "/dashboard/locacao/imoveis",
            dynamicChildren: {
              id: {
                name: "Detalhes do Imóvel",
                isDynamic: true
              }
            }
          },
          proprietarios: {
            name: "Proprietários",
            href: "/dashboard/locacao/proprietarios",
            dynamicChildren: {
              id: {
                name: "Proprietário",
                isDynamic: true
              }
            }
          },
        }
      },
    }
  }
};