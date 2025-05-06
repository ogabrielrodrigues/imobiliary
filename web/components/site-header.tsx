"use client"

import { Sidebar } from "lucide-react"
import React from 'react'

import { SearchForm } from "@/components/search-form"
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb"
import { Button } from "@/components/ui/button"
import { Separator } from "@/components/ui/separator"
import { useSidebar } from "@/components/ui/sidebar"
import { generateBreadcrumbPath, routes } from "@/lib/routes"
import { usePathname } from "next/navigation"

export function SiteHeader(): React.ReactElement {
  const { toggleSidebar } = useSidebar();
  const pathname = usePathname();

  const breadcrumbPath = generateBreadcrumbPath({
    pathname, 
    routeStructure: routes
  });

  return (
    <header className="bg-background sticky top-0 z-50 flex w-full items-center border-b">
      <div className="flex h-(--header-height) w-full items-center gap-2 pl-4 pr-8">
        <Button
          className="h-8 w-8"
          variant="ghost"
          size="icon"
          onClick={toggleSidebar}
        >
          <Sidebar />
        </Button>
        <Separator orientation="vertical" className="mr-2 h-4" />
        <Breadcrumb className="hidden sm:block">
          <BreadcrumbList>
            {breadcrumbPath.map((item, index) => (
              <React.Fragment key={`breadcrumb-${item.href}-${index}`}>
                <BreadcrumbItem>
                  {item.isLast ? (
                    <BreadcrumbPage>{item.name}</BreadcrumbPage>
                  ) : (
                    <BreadcrumbLink href={item.href}>{item.name}</BreadcrumbLink>
                  )}
                </BreadcrumbItem>
                {!item.isLast && <BreadcrumbSeparator />}
              </React.Fragment>
            ))}
          </BreadcrumbList>
        </Breadcrumb>
        <SearchForm className="w-full sm:ml-auto sm:w-1/6" />
      </div>
    </header>
  );
}

// "use client"

// import React from 'react'
// import { Sidebar } from "lucide-react"

// import { SearchForm } from "@/components/search-form"
// import {
//   Breadcrumb,
//   BreadcrumbItem,
//   BreadcrumbLink,
//   BreadcrumbList,
//   BreadcrumbPage,
//   BreadcrumbSeparator,
// } from "@/components/ui/breadcrumb"
// import { Button } from "@/components/ui/button"
// import { Separator } from "@/components/ui/separator"
// import { useSidebar } from "@/components/ui/sidebar"
// import { usePathname } from "next/navigation"

// /**
//  * Interface para os parâmetros dinâmicos em rotas
//  */
// interface DynamicParam {
//   name: string;
//   isDynamic: boolean;
//   children?: RouteStructure;
// }

// /**
//  * Interface para as rotas na estrutura de navegação
//  */
// interface RouteNode {
//   name: string;
//   href: string;
//   children?: RouteStructure;
//   dynamicChildren?: Record<string, DynamicParam>;
// }

// /**
//  * Tipo para a estrutura de rotas
//  */
// type RouteStructure = Record<string, RouteNode>;

// /**
//  * Interface para item do breadcrumb
//  */
// interface BreadcrumbItem {
//   name: string;
//   href: string;
//   isLast: boolean;
//   dynamicId?: string;
// }

// /**
//  * Estrutura de dados hierárquica para as rotas
//  */
// const routeStructure: RouteStructure = {
//   dashboard: {
//     name: "Visão Geral",
//     href: "/dashboard",
//     children: {
//       conta: {
//         name: "Conta",
//         href: "/dashboard/conta",
//       },
//       assinatura: {
//         name: "Assinatura",
//         href: "/dashboard/assinatura",
//       },
//       locacao: {
//         name: "Alugueis",
//         href: "/dashboard/locacao",
//         children: {
//           contratos: {
//             name: "Contratos",
//             href: "/dashboard/locacao/contratos",
//             dynamicChildren: {
//               id: {
//                 name: "Detalhes do Contrato",
//                 isDynamic: true
//               }
//             }
//           },
//           inquilinos: {
//             name: "Inquilinos",
//             href: "/dashboard/locacao/inquilinos",
//             dynamicChildren: {
//               id: {
//                 name: "Perfil do Inquilino", 
//                 isDynamic: true
//               }
//             }
//           },
//           imoveis: {
//             name: "Imóveis",
//             href: "/dashboard/locacao/imoveis",
//             dynamicChildren: {
//               id: {
//                 name: "Detalhes do Imóvel",
//                 isDynamic: true
//               }
//             }
//           },
//           proprietarios: {
//             name: "Proprietários",
//             href: "/dashboard/locacao/proprietarios",
//             dynamicChildren: {
//               id: {
//                 name: "Perfil do Proprietário",
//                 isDynamic: true
//               }
//             }
//           },
//           vistorias: {
//             name: "Vistorias",
//             href: "/dashboard/locacao/vistorias",
//           },
//           relatorios: {
//             name: "Relatórios",
//             href: "/dashboard/locacao/relatorios",
//           }
//         }
//       },
//     }
//   }
// };

// /**
//  * Componente de cabeçalho do site com breadcrumb dinâmico
//  */
// export function SiteHeader(): React.ReactElement {
//   const { toggleSidebar } = useSidebar();
//   const pathname = usePathname();

//   /**
//    * Verifica se um segmento parece ser um ID dinâmico
//    * @param segment - O segmento de URL a ser verificado
//    * @returns True se parece ser um ID
//    */
//   const isDynamicSegment = (segment: string): boolean => {
//     return (
//       // Verifica se é um número
//       /^\d+$/.test(segment) ||
//       // Verifica se é um UUID padrão
//       /^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i.test(segment) ||
//       // Verifica se é um código alfanumérico (como SKUs do produto)
//       /^[a-z0-9_-]{4,}$/i.test(segment)
//     );
//   };

//   /**
//    * Capitaliza a primeira letra de uma string
//    * @param text - Texto a ser capitalizado
//    * @returns Texto com primeira letra em maiúscula
//    */
//   const capitalizeFirstLetter = (text: string): string => {
//     return text.charAt(0).toUpperCase() + text.slice(1);
//   };

//   /**
//    * Gera o caminho completo do breadcrumb com base no pathname atual
//    * @returns Array de objetos breadcrumb
//    */
//   const generateBreadcrumbPath = (): BreadcrumbItem[] => {
//     // Segurança para pathname indefinido
//     if (!pathname) return [];
    
//     const pathSegments = pathname.split('/').filter(Boolean);
//     if (pathSegments.length === 0) return [];
    
//     const breadcrumbItems: BreadcrumbItem[] = [];
    
//     // Começamos verificando se temos o nó raiz 'dashboard'
//     if (!routeStructure.dashboard) {
//       return [];
//     }
    
//     let currentPath = '';
//     let currentStructure: RouteStructure | undefined = routeStructure;
//     let dynamicParent: RouteNode | null = null;
    
//     for (let i = 0; i < pathSegments.length; i++) {
//       const segment = pathSegments[i];
//       currentPath += `/${segment}`;
      
//       // Caso 1: Estamos no primeiro nível (dashboard)
//       if (i === 0 && segment === 'dashboard') {
//         const dashboardNode = currentStructure.dashboard;
        
//         breadcrumbItems.push({
//           name: dashboardNode.name,
//           href: dashboardNode.href,
//           isLast: pathSegments.length === 1
//         });
        
//         // Atualiza para o próximo nível se existir
//         currentStructure = dashboardNode.children;
//       } 
//       // Caso 2: Níveis subsequentes - rotas estáticas definidas
//       else if (currentStructure && segment in currentStructure) {
//         const currentNode = currentStructure[segment];
        
//         breadcrumbItems.push({
//           name: currentNode.name,
//           href: currentNode.href,
//           isLast: i === pathSegments.length - 1
//         });
        
//         // Armazena a estrutura atual como potencial pai de parâmetros dinâmicos
//         dynamicParent = currentNode;
        
//         // Atualiza para o próximo nível se existir
//         currentStructure = currentNode.children;
//       }
//       // Caso 3: Verifica se é um segmento dinâmico (como um ID)
//       else if (
//         isDynamicSegment(segment) && 
//         dynamicParent && 
//         dynamicParent.dynamicChildren && 
//         'id' in dynamicParent.dynamicChildren
//       ) {
//         const dynamicConfig = dynamicParent.dynamicChildren.id;
        
//         if (dynamicConfig) {
//           breadcrumbItems.push({
//             name: dynamicConfig.name,
//             href: currentPath,
//             isLast: i === pathSegments.length - 1,
//             dynamicId: segment
//           });
          
//           // Atualiza para possíveis rotas ainda mais profundas após o ID
//           currentStructure = dynamicConfig.children;
//         }
//       }
//       // Caso 4: Segmentos não mapeados (fallback)
//       else {
//         breadcrumbItems.push({
//           name: capitalizeFirstLetter(segment),
//           href: currentPath,
//           isLast: i === pathSegments.length - 1
//         });
        
//         // Reset da estrutura pois estamos em área não mapeada
//         currentStructure = undefined;
//       }
//     }
    
//     return breadcrumbItems;
//   };

//   // Gere o caminho do breadcrumb
//   const breadcrumbPath = generateBreadcrumbPath();

//   return (
//     <header className="bg-background sticky top-0 z-50 flex w-full items-center border-b">
//       <div className="flex h-(--header-height) w-full items-center gap-2 pl-4 pr-8">
//         <Button
//           className="h-8 w-8"
//           variant="ghost"
//           size="icon"
//           onClick={toggleSidebar}
//         >
//           <Sidebar />
//         </Button>
//         <Separator orientation="vertical" className="mr-2 h-4" />
//         <Breadcrumb className="hidden sm:block">
//           <BreadcrumbList>
//             {breadcrumbPath.map((item, index) => (
//               <React.Fragment key={`breadcrumb-${item.href}-${index}`}>
//                 <BreadcrumbItem>
//                   {item.isLast ? (
//                     <BreadcrumbPage>{item.name}</BreadcrumbPage>
//                   ) : (
//                     <BreadcrumbLink href={item.href}>{item.name}</BreadcrumbLink>
//                   )}
//                 </BreadcrumbItem>
//                 {!item.isLast && <BreadcrumbSeparator />}
//               </React.Fragment>
//             ))}
//           </BreadcrumbList>
//         </Breadcrumb>
//         <SearchForm className="w-full sm:ml-auto sm:w-1/6" />
//       </div>
//     </header>
//   );
// }
// // "use client"

// // import { Sidebar } from "lucide-react"

// // import { SearchForm } from "@/components/search-form"
// // import {
// //   Breadcrumb,
// //   BreadcrumbItem,
// //   BreadcrumbLink,
// //   BreadcrumbList,
// //   BreadcrumbPage,
// //   BreadcrumbSeparator,
// // } from "@/components/ui/breadcrumb"
// // import { Button } from "@/components/ui/button"
// // import { Separator } from "@/components/ui/separator"
// // import { useSidebar } from "@/components/ui/sidebar"
// // import { usePathname } from "next/navigation"
// // import { Fragment } from "react"

// // type DynamicParam = {
// //   name: string;
// //   isDynamic: boolean;
// //   children?: RouteStructure;
// // }


// // type RouteNode = {
// //   name: string;
// //   href: string;
// //   children?: RouteStructure;
// //   dynamicChildren?: Record<string, DynamicParam>;
// // }

// // type RouteStructure = Record<string, RouteNode>;

// // type BreadcrumbItem = {
// //   name: string;
// //   href: string;
// //   isLast: boolean;
// //   dynamicId?: string;
// // }

// // const routeStructure: RouteStructure = {
// //   dashboard: {
// //     name: "Visão Geral",
// //     href: "/dashboard",
// //     children: {
// //       conta: {
// //         name: "Conta",
// //         href: "/dashboard/conta",
// //       },
// //       assinatura: {
// //         name: "Assinatura",
// //         href: "/dashboard/assinatura",
// //       },
// //       locacao: {
// //         name: "Alugueis",
// //         href: "/dashboard/locacao",
// //         children: {
// //           contratos: {
// //             name: "Contratos",
// //             href: "/dashboard/locacao/contratos",
// //             dynamicChildren: {
// //               id: {
// //                 name: "Detalhes do Contrato",
// //                 isDynamic: true
// //               }
// //             }
// //           },
// //           inquilinos: {
// //             name: "Inquilinos",
// //             href: "/dashboard/locacao/inquilinos",
// //             dynamicChildren: {
// //               id: {
// //                 name: "Perfil do Inquilino",
// //                 isDynamic: true
// //               }
// //             }
// //           },
// //           imoveis: {
// //             name: "Imóveis",
// //             href: "/dashboard/locacao/imoveis",
// //             dynamicChildren: {
// //               id: {
// //                 name: "Detalhes do Imóvel",
// //                 isDynamic: true
// //               }
// //             }
// //           },
// //           proprietarios: {
// //             name: "Proprietários",
// //             href: "/dashboard/locacao/proprietarios",
// //             dynamicChildren: {
// //               id: {
// //                 name: "Perfil do Proprietário",
// //                 isDynamic: true
// //               }
// //             }
// //           },
// //           vistorias: {
// //             name: "Vistorias",
// //             href: "/dashboard/locacao/vistorias",
// //           },
// //           relatorios: {
// //             name: "Relatórios",
// //             href: "/dashboard/locacao/relatorios",
// //           }
// //         }
// //       },
// //     }
// //   }
// // };

// // export function SiteHeader() {
// //   const { toggleSidebar } = useSidebar();
// //   const pathname = usePathname();

// //   const isDynamicSegment = (segment: string): boolean => {
// //     return (
// //       /^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i.test(segment)
// //     );
// //   };

// //   const generateBreadcrumbPath = (): BreadcrumbItem[] => {
// //     const pathSegments = pathname.split('/').filter(Boolean);

// //     if (pathSegments.length === 0) return [];

// //     const breadcrumbItems: BreadcrumbItem[] = [];

// //     let currentStructure: RouteStructure = routeStructure;
// //     let currentPath = '';
// //     let dynamicParent: RouteNode | null = null;

// //     for (let i = 0; i < pathSegments.length; i++) {
// //       const segment = pathSegments[i];
// //       currentPath += `/${segment}`;

// //       // if (currentStructure === undefined) {
// //       //   return []
// //       // }

// //       if (i === 0 && segment === 'dashboard' && currentStructure.dashboard) {
// //         breadcrumbItems.push({
// //           name: currentStructure.dashboard.name,
// //           href: currentStructure.dashboard.href,
// //           isLast: pathSegments.length === 1
// //         });

// //         currentStructure = currentStructure.dashboard.children!;
// //       } else if (currentStructure && currentStructure[segment]) {
// //         breadcrumbItems.push({
// //           name: currentStructure[segment].name,
// //           href: currentStructure[segment].href,
// //           isLast: i === pathSegments.length - 1
// //         });

// //         dynamicParent = currentStructure[segment];

// //         if (!currentStructure[segment].children) return []

// //         currentStructure = currentStructure[segment].children;
// //       } else if (isDynamicSegment(segment) && dynamicParent?.dynamicChildren?.id) {
// //         const dynamicConfig = dynamicParent.dynamicChildren.id;

// //         breadcrumbItems.push({
// //           name: dynamicConfig.name,

// //           href: currentPath,
// //           isLast: i === pathSegments.length - 1,

// //           dynamicId: segment
// //         });

// //         if (!dynamicConfig.children) return []
// //         currentStructure = dynamicConfig.children;
// //       } else {
// //         breadcrumbItems.push({
// //           name: capitalizeFirstLetter(segment),
// //           href: currentPath,
// //           isLast: i === pathSegments.length - 1
// //         });
// //       }
// //     }

// //     return breadcrumbItems;
// //   };

// //   const capitalizeFirstLetter = (text: string): string => {
// //     return text.charAt(0).toUpperCase() + text.slice(1);
// //   };

// //   const breadcrumbPath = generateBreadcrumbPath();

// //   console.log(breadcrumbPath)

// //   return (
// //     <header className="bg-background sticky top-0 z-50 flex w-full items-center border-b">
// //       <div className="flex h-(--header-height) w-full items-center gap-2 pl-4 pr-8">
// //         <Button
// //           className="h-8 w-8"
// //           variant="ghost"
// //           size="icon"
// //           onClick={toggleSidebar}
// //         >
// //           <Sidebar />
// //         </Button>
// //         <Separator orientation="vertical" className="mr-2 h-4" />
// //         <Breadcrumb className="hidden sm:block">
// //           <BreadcrumbList>
// //             <BreadcrumbList>
// //               {breadcrumbPath.map((item, index) => (
// //                 <Fragment key={item.href}>
// //                   <BreadcrumbItem>
// //                     {item.isLast ? (
// //                       <BreadcrumbPage>{item.name}</BreadcrumbPage>
// //                     ) : (
// //                       <BreadcrumbLink href={item.href}>{item.name}</BreadcrumbLink>
// //                     )}
// //                   </BreadcrumbItem>
// //                   {!item.isLast && <BreadcrumbSeparator />}
// //                 </Fragment>
// //               ))}
// //             </BreadcrumbList>
// //             {/* <BreadcrumbItem>
// //               <BreadcrumbLink href="/dashboard">
// //                 Dashboard
// //               </BreadcrumbLink>
// //             </BreadcrumbItem>
// //             <BreadcrumbSeparator />
// //             {currentPage?.menu ? (
// //               <>
// //                 <BreadcrumbItem>
// //                   <BreadcrumbLink href={currentPage.menu.href}>
// //                     {currentPage.menu.name}
// //                   </BreadcrumbLink>
// //                 </BreadcrumbItem>
// //                 <BreadcrumbSeparator />
// //                 <BreadcrumbItem>
// //                   <BreadcrumbPage>{currentPage.name}</BreadcrumbPage>
// //                 </BreadcrumbItem>
// //               </>
// //             ) : (
// //               <BreadcrumbItem>
// //                 <BreadcrumbPage>{currentPage?.name}</BreadcrumbPage>
// //               </BreadcrumbItem>
// //             )} */}
// //           </BreadcrumbList>
// //         </Breadcrumb>
// //         <SearchForm className="w-full sm:ml-auto sm:w-1/6" />
// //       </div>
// //     </header>
// //   )
// // }
