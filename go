##!/bin/bash

NAM=`basename "$0"`
CWD=`pwd`
GOSUB="/usr/local/sd/syschk/lib"

. "$GOSUB""/f_go"

: ${PKG:="gotest"}
: ${VER:="0.0.0"}

F_GO_TMPFS=1G
#F_GO_DEL=src

CFG=
PFX=

fcp_usage ()
{
cat<<EOF
${NAM}: [ --help | -h ]
${NAM}: [ arc ] [-d ]
${NAM}: [ del ] [ all | obj | src ]

'configure'..
${NAM}: [ agen ] Autotool stuff (usually requires manual intervention).
${NAM}: [ acfg ] cd $OBJ && ../$SRC/configure.
${NAM}: [ amak ] cd $OBJ && make -jN.
${NAM}: [ ains ] cd $OBJ && make install (to $PFX target).
${NAM}: [ arem ] cd $OBJ && make uninstall (often doesn't work).

'cmake'..
${NAM}: [ ccfg ] cd $OBJ && cmake \$@ \$CFG ../$SRC
${NAM}: [ cmak ] cd $OBJ && cmake --build . \$@
${NAM}: [ cins ] cd $OBJ && cmake --install . \$@

${NAM}: [ all ]  User needs to modify this.

arc -d:
Unpack $PKG-$VER.tar.{gz,bz2,xz}

del all: Invoke "del obj" and "del src"
del obj: Remove $OBJ
del src: Remove $SRC

all: User modifies this to be able to unpack,build,install,delete build in
     one operation.

\$PFX is the install prefix (default /usr/local/toupper($PKG)/$VER)
and is currently: $PFX.

Don't mess with \$SRC,\$OBJ - they get set from \$PKG,\$VER.

F_GO_TMPFS specifies how large the SRC,OBJ tmpfs mounts will be. Comment out
to use the disk like normal. 32bit systems may fail after 4G but check inodes.

F_GO_DEL can be uncommented set to src or all to prevent go del [src|all] from
removing SRC,all - handy for those pesky in-tree builds or development.

You really ought to look in $SUB/f_go before fiddling!

Your configure options go in CFG after f_go_init in fcp_main, noting you don't
pass --prefix or INSTALL_PREFIX there: set PFX for that.

Should you uncomment 'f_go_time', ignore the warning about 'sdtimetool'. That
program contains propietory code currently so is not shipped atm. All it does
is show the elapsed time in an easy to view manner. Just use 'time' for now.

Ordinarily you'll be using one or other of these (autotool/cmake) methods.
See 'go.gcc' for an autotool example and 'go.db' for a cmake method.
EOF
}

fcp_arc ()
{
 f_go_arc "$@"
}

fcp_agen ()
{
 f_go_gen "$@"
}

fcp_acfg ()
{
 f_go_cfg "$@"
}

fcp_amak ()
{
 f_go_mak "$@"
}

fcp_ains ()
{
 f_go_ins "$@"
}

fcp_arem ()
{
 f_go_rem "$@"
}

fcp_ccfg ()
{
 f_go_ccfg "$@"
}

fcp_cmak ()
{
 f_go_cmak "$@"
}

fcp_cins ()
{
 f_go_cins "$@"
}

fcp_crem ()
{
 f_go_rem "$@"
}

fcp_del ()
{
 f_go_del "$@"
}

ARG="$1"
shift
fcp_main ()
{
 f_go_init
PFX="/tmp/${PKG}-${VER}"
CFG="
"

 f_go_time_b
 case "$ARG" in
	--help | -h)
	fcp_usage
	exit 0
	;;

	arc)
	fcp_arc "$@" "$SRC"
	;;

	agen)
	fcp_agen "$@"
	;;

	acfg)
	fcp_acfg "$@"
	;;

	amak)
	fcp_amak "$@"
	;;

	ains)
	fcp_ains "$@"
	;;

	arem)
	fcp_arem "$@"
	;;

	ccfg)
	fcp_ccfg "$@"
	;;

	cmak)
	fcp_cmak "$@"
	;;

	cins)
	fcp_cins "$@"
	;;

	crem)
	fcp_crem "$@"
	;;

	del)
	fcp_del "$@"
	;;

	cfg)
	fcp_acfg "$@"
#	fcp_ccfg "$@"
	;;

	mak)
	fcp_amak "$@"
#	fcp_cmak "$@"
	;;

	ins)
	fcp_ains "$@"
#	fcp_cins "$@"
	;;

	aall)
	fcp_arc -d "$SRC" || exit 1
	fcp_agen
	fcp_acfg || exit 1
	fcp_amak -j `f_go_bproc` || exit 1
	fcp_ains || exit 1
	"$PFX"/bin/"$PKG"
	fcp_arem
	fcp_del all
	;;

	call)
	fcp_arc -d "$SRC" || exit 1
	fcp_ccfg || exit 1
	fcp_cmak -j `f_go_bproc` || exit 1
	fcp_cins || exit 1
	"$PFX"/bin/"$PKG"
#	fcp_crem
	fcp_del all
	;;

	all)
	(
	./"$NAM" aall || exit 1
	./"$NAM" call || exit 1
	) || {
		echo -e "\n\n$NAM: -ERR: test build" 1>&2
		exit 1
	}
	echo -e "\n\n$NAM: +OKI: Test build"
	;;

	*)
	echo "Bad command!" 1>&2
	exit 1
	;;
 esac
 RETV=$?
 echo "($NAM: $0)[""$RETV"']'
 f_go_time_e
# f_go_time
 return $RETV
}

fcp_main "$@"
