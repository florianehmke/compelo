export function tokenForProjectIdExists(id: number): boolean {
  const token = localStorage.getItem('compelo-token');
  if (token) {
    const decoded = JSON.parse(atob(token.split('.')[1]));
    return decoded.projectId === id;
  }
  return false;
}
