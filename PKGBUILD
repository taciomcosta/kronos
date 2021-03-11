# Maintainer: Tacio Costa <taciomcosta@gmail.com>
pkgname=kronos
pkgver=0.2.0
pkgrel=1
pkgdesc="Cross-platform job scheduler for cronjobs"
arch=('any')
url="https://github.com/taciomcosta/$pkgname"
license=('MIT')
makedepends=('go')
source=("$url/releases/download/0.2.0/kronos-$pkgver-linux_amd64.tar.gz")
sha256sums=('bd471668d9b1ad2e656724db9a5b3a9246b9ed2e2722220d1b7b20b5de2a99e0')

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
