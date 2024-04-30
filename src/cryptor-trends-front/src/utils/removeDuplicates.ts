export function removeDuplicates<T>(lista: T[]): T[] {
  return lista.filter((item, index) => lista.indexOf(item) === index);
}
