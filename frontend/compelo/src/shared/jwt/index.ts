export function tokenForProjectIdExists(id: number): boolean {
  return projectIdFromToken() === id;
}

export function projectIdFromToken(): number {
  const token = loadToken();
  if (token) {
    const decoded = JSON.parse(atob(token.split('.')[1]));
    return decoded.projectId;
  }
  return -1;
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
