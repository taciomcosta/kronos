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
	cd "$pkgname-$pkgver"
  mkdir -p build/
}

build() {
  cd "$pkgname-$pkgver"
  mkdir -p /usr/local/var/kronos
  mkdir -p /usr/local/etc/kronos
	make release-linux
}

check() {
  make test-unit
}

package() {
  cd "$pkgname-$pkgver"
  install -Dm756 build/$pkgname "$pkgdir"/usr/bin/$pkgname
}
