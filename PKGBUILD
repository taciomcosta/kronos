# Maintainer: Tacio Costa <taciomcosta@gmail.com>
pkgname=kronos
pkgver=0.2.0
pkgrel=1
pkgdesc="Cross-platform job scheduler for cronjobs"
arch=('x86_64')
url="https://github.com/taciomcosta/$pkgname"
license=('MIT')
makedepends=('go')
source=("$url/releases/download/0.2.0/kronos-$pkgver-linux_amd64.tar.gz")
sha256sums=('b9f738a6fffd75c669e7529711547f09a6d7b78ac50ac1481770d73b9ba8ff5d')

prepare() {
	cd "$srcdir"
  mkdir -p build/
}

build() {
  cd "$srcdir"
  sudo mkdir -p /usr/local/var/kronos
  sudo mkdir -p /usr/local/etc/kronos
}

package() {
  cd "$srcdir"
  install -Dm755 build/kronos "$pkgdir"/usr/bin/kronos
  install -Dm755 build/kronosd "$pkgdir"/usr/bin/kronosd
}
