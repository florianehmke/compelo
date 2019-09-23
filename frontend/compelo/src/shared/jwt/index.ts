export function tokenForProjectIdExists(id: number): boolean {
  const token = loadToken();
  if (token) {
    const decoded = JSON.parse(atob(token.split('.')[1]));
    return decoded.projectId === id;
  }
  return false;
}

export function loadToken(): string {
  return localStorage.getItem('compelo-token');
}

export function removeToken() {
  localStorage.removeItem('compelo-token');
}

export function storeToken(token: string) {
  localStorage.setItem('compelo-token', token);
}
