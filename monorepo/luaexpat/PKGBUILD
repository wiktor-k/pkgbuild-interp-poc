# Maintainer: Sergej Pupykin <pupykin.s+arch@gmail.com>
# Maintainer: Daurnimator <daurnimator@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Lucas Hermann Negri <kkndrox@gmail.com>

pkgbase=luaexpat
pkgname=(lua-expat lua51-expat lua52-expat lua53-expat lua54-expat)
pkgver=1.5.2
pkgrel=2
_rockrel=1
pkgdesc='SAX XML parser based on the Expat library'
arch=('x86_64')
url="https://github.com/lunarmodules/$pkgbase"
license=('MIT')
depends=(expat)
makedepends=(lua
             lua51
             lua52
             lua53
             lua54
             luarocks)
_archive="$pkgbase-$pkgver"
_rock="$_archive-$_rockrel.linux-$CARCH.rock"
_rockspec="$_archive-$_rockrel.rockspec"
source=("$url/archive/$pkgver/$_archive.tar.gz")
sha512sums=('a2a04b2bb76a10a8a0da077b53d22b6e84460f416486a4b006d3df4991cff011dc91a225ab25f994a2713632d562b13d4341530bbc7ede4fd660a6fce24eed3c')

build() {
	cd "$_archive"
	for LUAVER in 5.{1,2,3,4,5}; do
		luarocks --lua-version "$LUAVER" \
			CFLAGS="$CFLAGS -DLUA_C89_NUMBERS" \
			make --pack-binary-rock --deps-mode none -- "rockspecs/$_rockspec"
		install -Dm0644 -t "lua-$LUAVER/" "$_rock"
	done
}

_package() {
	cd "$_archive"
	depends+=("${pkgname%-*}")
	luarocks --lua-version "$1" --tree "$pkgdir/usr/" \
		install --deps-mode none --no-manifest -- "lua-$1/$_rock"
	install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE
}

package_lua-expat() {
	_package 5.5
}

package_lua54-expat() {
	_package 5.4
}

package_lua53-expat() {
	_package 5.3
}

package_lua52-expat() {
	_package 5.2
}

package_lua51-expat() {
	replaces=(luaexpat)
	conflicts=(luaexpat)
	_package 5.1
}
