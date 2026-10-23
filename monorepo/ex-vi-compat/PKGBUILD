# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=ex-vi-compat
pkgver=3
pkgrel=1
pkgdesc="The ex and vi commands based on vim's compatibility modes"
arch=(any)
url="https://gitlab.archlinux.org/archlinux/packaging/packages/ex-vi-compat"
license=(0BSD)
depends=(
  bash
  vim
)
provides=(
  edit
  ex
  vedit
  vi
  view
)
conflicts=(vi)
replaces=(vi)
source=(
  LICENSE
  vi.sh
)
sha512sums=('a33658d9271e5c537ccd41bf540b463ad2a5eca4a060c80486ff42a736f0aa042d10436e7177c34d792177cb11285243dee1f31c4df54fb0bfaabbc306406930'
            'a316d8622a9358c4b0e3b014bbf76d779c8d0d8c5da843f48bf4a72427f7e7947203dd6f6543a652e82a86c3569de07f4ae442a396e45914481beca2d9ef031f')
b2sums=('a29664104e1ee73ca0aee1d633e9095d92a57c92787f8d8740bdb7211ba3205782ed8677f539bdb8cae3dd75a3694be3132e185fa3fc4b3f401e1f88eb776101'
        'cac444af20f5a304e32f3499e911fa80a06f49e59076e0c36a0534951cb1d456bc66f82d621fe90b675d18a310e8634ca18b681adae8e3cf7152336385399fb5')

package() {
  local cmd

  install -vDm 755 vi.sh "$pkgdir/usr/bin/vi"
  for cmd in {edit,ex,vedit,view} ; do
    ln -sv vim "$pkgdir/usr/bin/$cmd"
  done
  install -vDm 644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/0BSD.txt"
}
